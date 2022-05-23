package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
)

func main() {
	data := getInput("./test.txt")	
}

func selector(lines []string) {
	// Select between addition, exploding or splitting
	finalNumber := string
	for _, nextNumber := range lines {
		// Add next number
		add()

		for {
			var clear1, clear2 bool

			// First, check for nested pairs to explode. Explode if nested in 4 pairs
			nested := lookForNestedPairs(finalNumber)
			if nested == 0 {
				clear1 = true
			}

			// Second, check for splits. Split if any number > 10
			location := lookForBigNumbers()
			if location == 0 {
				clear2 = true
			}
			if clear1 == true && clear2 == true {
				break
			}
		}
	}

}

func add() {

}

func explode() {

}

func split() {

}

func lookForBigNumbers(snailfishNumber string) int {
	// Parse string to []byte to get accepted by regexp.MustCompile
	content := []byte(snailfishNumber)

	// Find all numbers bigger than 9
	pattern := regexp.MustCompile(`[1-9][0-9]`)
	loc := pattern.FindIndex(content)
	fmt.Println(loc)
	fmt.Println(string(content[loc[0]:loc[1]]))
}

func lookForNestedPairs(snailfishNumber string) int {
	// Check for nested pairs
	// Returns the location of the left number needing to explode
	leftBracket := 0
	for i, ch := range snailfishNumber {
		if ch == '[' {
			leftBracket++
		}

		if leftBracket == 5 {
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