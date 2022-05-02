package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	polymer, rules := getPuzzleInput("./14.test")
	fmt.Printf("Template:\t%s\n", polymer)

	instructions := make(map[int][]string, 0)
	for i, rule := range rules {
		instructions[i] = strings.Split(rule, " -> ")
	}

	// Pair insertion. [N (N] [C) B] -> [N C (N] B [C) H B]
	for steps := 0; steps <= 40; steps++ {
		polymer = insertion(polymer, instructions)
		fmt.Printf("After step %d:\t%d\n", steps, len(polymer))
		if steps == 9 {
			result := countElements(polymer)
			fmt.Printf("Part 1 -> %d\n", result)
		}

		if steps == 40 {
			result := countElements(polymer)
			fmt.Printf("Part 2 -> %d\n", result)
		}
	}
	
}

func countElements(polymer string) int {
	occurrences := make(map[string]int)
	
	for _, ch := range polymer {
		occurrences[string(ch)]++
	}

	fmt.Println(occurrences)

	var mostCommon, leastCommon int

	for _, v := range occurrences {
		if leastCommon == 0 {
			leastCommon = v
		}

		if v > mostCommon {
			mostCommon = v
		} else if v < leastCommon {
			leastCommon = v
		}
	}

	return mostCommon - leastCommon
}

func insertion(template string, instructions map[int][]string) string {
	separated := make([]string, 0)
	finalString := make([]string, 0)
	
	var newString string
	
	for i, ch := range template {
		if i+1 < len(template) {
			newString = string(ch) + string(template[i+1])
		}
		
		separated = append(separated, newString)
		newString = ""
	}

	// for i, rule := range rules {
	// 	instructions[i] = strings.Split(rule, " -> ")
	// }

	// Pair insertion. [N (N] [C) B] -> [N C (N] B [C) H B]
	for i, str := range separated {

		for _, instruction := range instructions {

			if str == instruction[0] {
				myStr := string(str[0]) + instruction[1]

				if i == len(separated)-2 {
					myStr += string(str[1])
				}

				finalString = append(finalString, myStr)
			}
		}

		
	}

	return strings.Join(finalString, "") 
}

func getPuzzleInput(filename string) (string, []string) {
	bytes, _ := ioutil.ReadFile(filename)

	dirtyLines := strings.Split(string(bytes), "\n")
	lines := make([]string, 0)
	for _, line := range dirtyLines {
		lines = append(lines, strings.TrimSuffix(line, "\r"))
	}

	template := lines[0]
	rules := lines[2:]

	return template, rules
}

