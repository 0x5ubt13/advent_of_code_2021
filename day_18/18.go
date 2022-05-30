package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	selector(getInput("./test.txt"))	
}

func selector(lines []string) {
	// Select between addition or reduction (exploding or splitting)
	var finalNumber string

	// loop over the new numbers
	for i, nextNumber := range lines {
		if i == 0 {
			finalNumber = nextNumber
			fmt.Println(finalNumber)

			continue
		}

		// Add next number
		finalNumber = add(finalNumber, nextNumber)
		fmt.Println(finalNumber)

		// Reduce number
		finalNumber = reduce(finalNumber)

		// Repeat
	}
	// finalNumber = lines[0]
	finalNumber = reduce(finalNumber)

	fmt.Println(finalNumber)
}

func add(finalNumber, nextNumber string) string {
	return fmt.Sprintf("[%s,%s]", finalNumber, nextNumber)
}

func reduce(finalNumber string) string {
	// Loop through the number until no more reductions can be done

	// Add a pointer to keep track of the position
	// var pointer int = len(finalNumber)

	for {
		var clear1, clear2 bool

		// First, check for nested pairs to explode. Explode pair if pair is nested inside 4 pairs
		nestedLocationLeft, _ := lookForNestedPairs(finalNumber)
		// pointer = nestedLocationRight
		if nestedLocationLeft == 0 {
			clear1 = true
		} else {
			fmt.Println(finalNumber)
			// finalNumber = explode(finalNumber, nestedLocationLeft, nestedLocationRight)
			
			left, right := extractNestedSlice(finalNumber)
			bigNumLocation := lookForBigNumbers(finalNumber[left:right])
			if len(bigNumLocation) != 0 {
				bigNumLocation := lookForBigNumbers(finalNumber)
				finalNumber = split(finalNumber, bigNumLocation)
			} 
			nestedLocationLeft, nestedLocationRight := lookForNestedPairs(finalNumber)
			// finalNumber = explode(finalNumber, left, right)
			finalNumber = explode(finalNumber, nestedLocationLeft, nestedLocationRight)
		}

		// if clear1 == true {
		// 	break // debug break
		// }


		// Second, check for splits. Split if any number > 10
		bigNumLocation := lookForBigNumbers(finalNumber)
		if len(bigNumLocation) == 0 {
			clear2 = true
		} else {
			// TODO: if found big num what
			fmt.Println(finalNumber)
			finalNumber = split(finalNumber, bigNumLocation)
			// nestedLocationLeft, nestedLocationRight := lookForNestedPairs(finalNumber)
			// // pointer = nestedLocationRight
			// if nestedLocationLeft == 0 {
			// 	clear1 = true
			// } else {
			// 	fmt.Println(finalNumber)
			// 	finalNumber = explode(finalNumber, nestedLocationLeft, nestedLocationRight)
			// }
		}
		

		if clear1 == true && clear2 == true {
			break
		}
	}

	return finalNumber
}

func explode(finalNumber string, nestedLocationLeft, nestedLocationRight int) string {
	// nested left and nested right are the locations of the numbers to "explode"
	// nested left -1 is the left bracket, right + 1 the right one
	fmt.Println(finalNumber, nestedLocationLeft, nestedLocationRight, "slice:", finalNumber[nestedLocationLeft:nestedLocationRight])

	// Transform to nums to sum with the next numbers on the sides, if any
	explodingLeft, err := strconv.Atoi(string(finalNumber[nestedLocationLeft:nestedLocationRight][1])); if err != nil{panic(err)}
	explodingRight, err := strconv.Atoi(string(finalNumber[nestedLocationLeft:nestedLocationRight][3])); if err != nil{panic(err)}
	
	fmt.Printf("Extracted: %v\n", finalNumber[nestedLocationLeft:nestedLocationRight])
	fmt.Printf("Numbers to work with: left: %d right: %d\n", explodingLeft, explodingRight)

	explodedLeft := finalNumber[:nestedLocationLeft]
	explodedRight := finalNumber[nestedLocationRight:]

	// [[[[  [1,1],[2,2]],[3,3]],[4,4]],[5,5]]
	// [[[[   0   ,[3,2]],[3,3]],[4,4]],[5,5]]
	// [[[[   3   ,     ],[5,3]],[4,4]],[5,5]]
	// [[[[   3   ,    0],[5,3]],[4,4]],[5,5]]

	// [[[[  [4,3] ,4  ],4],[7,[[8,4],9]]],[1,1]]
	// [[[[   0,    7  ],4],[7,[[8,4],9]]],[1,1]]
	explodedMiddle := ""
	
	// If no number left, plant a 0, else add the rightmost
	checkNumbersLeft := lookForNumbers(explodedLeft)
	if len(checkNumbersLeft) == 0 {
		explodedMiddle += "0"
	} else {
		// if not adjacent, add a 0 in the exploded pair 
		if explodedLeft[len(explodedLeft)-2] > 57 || explodedLeft[len(explodedLeft)-2] < 48 && explodedMiddle != "0" {
			explodedMiddle += "0"
		}
		// find the rightmost, convert to int, add exploding one, convert back to str and sub in str
		target, err := strconv.Atoi(string(explodedLeft[checkNumbersLeft[len(checkNumbersLeft)-1][0]])); if err != nil{panic(err)}
		target += explodingLeft
		explodedLeft = explodedLeft[:checkNumbersLeft[len(checkNumbersLeft)-1][0]] + strconv.Itoa(target) + explodedLeft[checkNumbersLeft[len(checkNumbersLeft)-1][0]+1:]
	}

	checkNumbersRight := lookForNumbers(explodedRight)
		// If no number right, plant a 0
		if len(checkNumbersRight) == 0 {
			explodedMiddle += "0"
		} else {
			// if not adjacent, add a 0 in the exploded pair 
			fmt.Println(explodedRight[0])
			if explodedRight[0] > 57 || explodedRight[0] < 48 && explodedMiddle != "0"{
				explodedMiddle += "0"
			}
			// find the leftmost, convert to int, add exploding one, convert back to string and substitute in the string
			target, err := strconv.Atoi(string(explodedRight[checkNumbersRight[0][0]])); if err != nil{panic(err)}
			target += explodingRight
			explodedRight = explodedRight[:checkNumbersRight[0][0]] + strconv.Itoa(target) + explodedRight[checkNumbersRight[0][0]+1:]
		}

	fmt.Println(explodedLeft + "/" + explodedMiddle + "/" + explodedRight)
	explodedNumber := explodedLeft + explodedMiddle + explodedRight 
 
	return explodedNumber
}

func split(finalNumber string, bigNumLocation []int) string {
	// bigNumLocation indicates slice to cut
	splittingLeft := finalNumber[:bigNumLocation[0]]
	splittingRight := finalNumber[bigNumLocation[1]:]
	splittingNum, err := strconv.Atoi(finalNumber[bigNumLocation[0]:bigNumLocation[1]]); if err!=nil{panic(err)}
	fmt.Println("Extracted:", splittingNum)
	fmt.Println(splittingLeft + "| |" + splittingRight)
	var leftSplit float64
	if splittingNum % 2 != 0 {
		leftSplit = math.Trunc(float64(splittingNum) / 2)
		rightSplit := leftSplit + 1
	
		return fmt.Sprintf("%v[%v,%v]%v", splittingLeft, int64(leftSplit), int64(rightSplit), splittingRight)
	} else {
		leftSplit = float64(splittingNum) / 2 
		
		return fmt.Sprintf("%v[%v,%v]%v", splittingLeft, int64(leftSplit), int64(leftSplit), splittingRight) 
	}
}

func extractNestedSlice(snailfishNumber string) (int, int) {
	// work on the level of the nested pair to avoid crashing the program
	begin := 0
	nesting := 0
	
	for i, ch := range snailfishNumber {
		if ch == '[' {
			begin++
		} 

		if ch == ']' {
			begin--
		}

		if begin == 5 {
			for j, ch2 := range snailfishNumber[i:] {
				if ch2 == '[' {
					nesting++
				}
				if ch2 == ']' {
					if nesting == 0 {
						return i, i+j
					}
					nesting--
				}
			}
		}
	}

	return 0, 0
}

func lookForNumbers(snailfishNumber string) [][]int {
	// Parse string to []byte to get accepted by regexp.MustCompile
	content := []byte(snailfishNumber)

	// RegExp to find all the numbers inside the slice of the number
	pattern := regexp.MustCompile(`[0-9]`)
	loc := pattern.FindAllIndex(content, -1)

	return loc
}

func lookForBigNumbers(snailfishNumber string) []int {
	// Parse string to []byte to get accepted by regexp.MustCompile
	content := []byte(snailfishNumber)

	// RegExp to find the first number bigger than 9
	pattern := regexp.MustCompile(`[1-9][0-9]`)
	loc := pattern.FindIndex(content)
	// fmt.Println(loc)
	// fmt.Println(string(content[loc[0]:loc[1]]))

	return loc
}

func lookForNestedPairs(snailfishNumber string) (int, int) {
	// Check for nested pairs
	leftBracket := 0
	for i, ch := range snailfishNumber {
		if ch == '[' {
			leftBracket++
		}
		
		if leftBracket == 5 {
			// Returns the location of the left bracket of the pair needing to explode
			// and the location of the comma afterwards
			for {
				if snailfishNumber[i+1] == '[' || snailfishNumber[i+1] == ',' || snailfishNumber[i+3] == '[' || snailfishNumber[i+3] == ',' {
					i++
				} else {
					return i, i+5
				}
			}
		}

		if ch == ']' {
			leftBracket--
		}
	}

	return 0, 0
}

func getInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename); if err != nil { panic(err) }

	lines := make([]string, 0)
	for _, line := range strings.Split(string(bytes), "\n") {
		lines = append(lines, line)
	}

	return lines
}