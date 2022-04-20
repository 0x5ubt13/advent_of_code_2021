package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
)

func main() {
	data := getPuzzleInput("./10.in")
	
	// ---------- Part 1 ---------- 
	valid := make(map[string]int)
	corrupted := make(map[string]int)
	const opening string = "([{<"
	const closing string = ")]}>"

	for _, line := range data {
		openingQueue := make([]string, 0)
		var corrupt bool

		for _, char := range line {
			// opening char
			if strings.Contains(opening, string(char)) == true {
				openingQueue = append(openingQueue, string(char))
				// fmt.Println("added", string(char))
				continue
			} 
			
			if strings.Contains(closing, string(char)) == false {
				continue
			}

			// closing char, make sure it matches openingQueue[0] 
			if openingQueue[len(openingQueue)-1] == "(" && char == ')' || 
				openingQueue[len(openingQueue)-1] == "[" && char == ']' ||
				openingQueue[len(openingQueue)-1] == "{" && char == '}' ||
				openingQueue[len(openingQueue)-1] == "<" && char == '>' {
				openingQueue = openingQueue[:len(openingQueue)-1]
			} else {
				// or else it gets corrupt
				corrupt = true
				corrupted[string(char)] += points(string(char))
				break
			}

		}

		if corrupt == false {
			valid[strings.Join(openingQueue, "")] = len(openingQueue)
		}
	}

	score := 0
	for _, v := range corrupted {
		score += v
	}

	// ---------- Part 2 ---------- 
	autocompletedScores := make([]int, 0, 0)
	for key, length := range valid {
		score := 0
		// work backwards
		for x := length-1; x > -1; x-- {
			score *= 5
			if key[x] == '(' {
				score += 1
			} else if key[x] == '[' {
				score += 2
			} else if key[x] == '{' {
				score += 3
			} else {
				score += 4
			}
		}
		
		autocompletedScores = append(autocompletedScores, score)
	}

	sort.Ints(autocompletedScores)
	fmt.Printf("Part 1 -> %d\nPart 2 -> %d", score, autocompletedScores[len(autocompletedScores)/2])
}

func points(char string) int {
	if char == ")" {
		return 3
	} else if char == "]" {
		return 57
	} else if char == "}" {
		return 1197
	} else if char == ">" {
		return 25137
	}

	return 0
}

func getPuzzleInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(string(bytes), "\n")
}