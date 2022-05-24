package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
)

func main() {
	selector(getInput("./test.txt"))	
}

func selector(lines []string) {
	var finalNumber string
	// Select between addition, exploding or splitting

	// loop over the new lines
	for i, nextNumber := range lines {
		if i == 0 {
			finalNumber = nextNumber
			continue
		}
		fmt.Println(finalNumber)

		// Add next number
		finalNumber = add(finalNumber, nextNumber)

		for {
			var clear1, clear2 bool

			// First, check for nested pairs to explode. Explode if nested in 4 pairs
			nestedLocation := lookForNestedPairs(finalNumber)
			if nestedLocation == 0 {
				clear1 = true
			} else {
				// You are here
				finalNumber = explode(finalNumber, nestedLocation)
			}


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
	}

	fmt.Println(finalNumber)
}

func add(finalNumber, nextNumber string) string {
	finalNumber = fmt.Sprintf("[%s,%s]", finalNumber, nextNumber)

	return finalNumber
}

// You are here:
func explode(finalNumber, nestedLocation) string {

}

func split() {

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