package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	selector(getInput("./test.txt"))	
}

func selector(lines []string) {
	var finalNumber string
	// Select between addition or reduction (exploding or splitting)

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

	fmt.Println(finalNumber)
}

func add(finalNumber, nextNumber string) string {
	return fmt.Sprintf("[%s,%s]", finalNumber, nextNumber)
}

func reduce(finalNumber string) string {
	// Loop through the number until no more reductions can be done
	for {
		var clear1, clear2 bool

		// First, check for nested pairs to explode. Explode pair if pair is nested inside 4 pairs
		nestedLocation := lookForNestedPairs(finalNumber)
		if nestedLocation == 0 {
			clear1 = true
		} else {
			// You are here
			finalNumber = explode(finalNumber, nestedLocation)
		}

		break

		// Second, check for splits. Split if any number > 10
		bigNumLocation := lookForBigNumbers(finalNumber)
		if len(bigNumLocation) == 0 {
			clear2 = true
		}
		// TODO: if found big num what

		if clear1 == true && clear2 == true {
			break
		}
	}

	return finalNumber
}

func explode(finalNumber string, nestedLocation int) string {
	// nested location is the location of the left number to explode
	explodedNumber := ""

	// [[[[[9 , 8],1],2],3],4] becomes [[[[ 0 ,  9 ],2],3],4]
	// [[[[[%s,%s],1],2],3],4] becomes [[[[ 0 , 8+1],2],3],4]

	// [[3,[2,[ 1 ,[  7 ,  3  ]]]],[  6 ,[5,[4,[3,2]]]]] becomes [[3,[2,[  8 ,  0  ]]],[  9  ,[5,[4,[3,2]]]]]
	// [[3,[2,[ 1 ,[ %s ,  %s ]]]],[  6 ,[5,[4,[3,2]]]]] becomes [[3,[2,[ 1+7,  0  ]]],[ 3+6 ,[5,[4,[3,2]]]]]
	//        ^ ^     ^ ,  ^          ^		   			 becomes        ^  ^ ,  ^         ^  ^
	// original first num , second num original        becomes   o    first num , second num original

	for i, ch := range finalNumber {
		if i == nestedLocation {
			explodingNumber, err := strconv.Atoi(string(ch)); if err != nil { panic(err) }

			// Explode left
			findNumbersLeft := lookForNumbers(finalNumber[:i])
			if len(findNumbersLeft) == 0 {
				explodedNumber = finalNumber[:i-1]
				explodedNumber += "0"
				explodedNumber += finalNumber[i:]
			} else {
				// Grab the rightmost one
				indexLeft := findNumbersLeft[len(findNumbersLeft)-1][0]
				candidateLeft, err := strconv.Atoi(string(finalNumber[indexLeft])); if err != nil{panic(err)}
				candidateLeft += explodingNumber

				explodedNumber = finalNumber[:i-1]
				explodedNumber += strconv.Itoa(candidateLeft)
				explodedNumber += finalNumber[i:]				
			}

			// Explode right
			// fmt.Println(finalNumber[i+1:])
			workingSlice := finalNumber[i+1:]
			findNumbersRight := lookForNumbers(finalNumber[i+1:])
			if len(findNumbersRight) == 0 {
				explodedNumber = finalNumber[:i-1]
				explodedNumber += "0"
				explodedNumber += finalNumber[i:]
			} else {
				// Grab the leftmost one
				indexRight := findNumbersRight[0][0]
				// fmt.Println(findNumbersRight)
				candidateRight, err := strconv.Atoi(string(workingSlice[indexRight])); if err != nil{panic(err)}
				candidateRight += explodingNumber

				explodedNumber = finalNumber[:i-1]
				explodedNumber += strconv.Itoa(candidateRight)
				explodedNumber += finalNumber[i:]	
			}

		}
	}

	return explodedNumber
}

func split() {

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

func lookForNestedPairs(snailfishNumber string) int {
	// Check for nested pairs
	leftBracket := 0
	for i, ch := range snailfishNumber {
		if ch == '[' {
			leftBracket++
		}
		
		if leftBracket == 5 {
			// Returns the location of the left number of the pair needing to explode
			return i+1
		}

		if ch == ']' {
			leftBracket--
		}
	}

	return 0
}

func getInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename); if err != nil { panic(err) }

	lines := make([]string, 0)
	for _, line := range strings.Split(string(bytes), "\n") {
		lines = append(lines, line)
	}

	return lines
}