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
		nestedLocationLeft, nestedLocationRight := lookForNestedPairs(finalNumber)
		if nestedLocationLeft == 0 {
			clear1 = true
		} else {
			// You are here
			finalNumber = explode(finalNumber, nestedLocationLeft, nestedLocationRight)
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

func explode(finalNumber string, nestedLocationLeft, nestedLocationRight int) string {
	// nested left and nested right are the locations of the numbers to "explode"
	// nested left -1 is the left bracket, right + 1 the right one

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
		// find the rightmost, convert to int and add exploding one

	}

	checkNumbersRight := lookForNumbers(explodedRight)
		// If no number right, plant a 0
		if len(checkNumbersRight) == 0 {
			explodedMiddle += "0"
		}

	explodedNumber := explodedLeft + " " + explodedMiddle + " " + explodedRight 
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
			return i, i+5
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