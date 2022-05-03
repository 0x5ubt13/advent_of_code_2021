package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	polymer, rules := getPuzzleInput("./14.test")
	fmt.Printf("Template:\t%s\n", polymer)

	instructions := make(map[string]*Output)
	for _, rule := range rules {
		var k, v string
		fmt.Sscanf(rule, "%s -> %s", &k, &v)

		instructions[k] = &Output{
			Pair1: 		string(k[0]) + v,
			Pair2: 		v + string(k[1]),
			NewElement: v,
		}
	}

	// instructions := make(map[int][]string, 0)
	// for i, rule := range rules {
	// 	instructions[i] = strings.Split(rule, " -> ")
	// }

	polymerMap := make(map[int]string)
	counter := 0
	for i, ch := range polymer {
		if i == len(polymer)-1 {
			polymerMap[counter] += string(ch)
		} else {
			polymerMap[counter] += string(ch) 
			if i % 2 == 1 {
				counter++
			}
		}
	}

	var result int

	// Pair insertion. [N (N] [C) B] -> [N C (N] B [C) H B]
	for steps := 0; steps <= 1; steps++ {
		polymerMap = insertion(polymerMap, instructions)
		fmt.Printf("After step %d:\t%d\n", steps, len(polymer))
		if steps == 9 {
			result = countElements(polymerMap)
			fmt.Printf("Part 1 -> %d\n", result)
		}
	}
}

// start
type Step struct {
	Pairs map[string]int
	Count map[string]int
}

type Output struct {
	Pair1 string
	Pair2 string
	NewElement string
}

func step(in *Step, rules map[string]Output) *Step {
	out := &Step{
		Pairs: make(map[string]int),
		Count: make(map[string]int),
	}

	// Copying the map
	for k, v := range in.Count {
		out.Count[k] = v
	}

	// Output -> rules
	for p, _ := range in.Pairs {
		t := rules[p]
		out.Pairs[t.Pair1]++
		out.Pairs[t.Pair2]++
		out.Count[t.NewElement]++
	}

	return out
}

func newStep(in string) *Step {
	s := &Step{
		Pairs: make(map[string]int),
		Count: make(map[string]int),
	}

	for i := 0; i < len(in)-1; i++ {
		pair := in[i : i+2]
		s.Pairs[pair]++
	}

	for i := 0; i < len(in)-1; i++ {
		pair := in[i : i+2]
		s.Count[string(in[i])]++
	}
}
//end

func countElements(polymerMap map[int]string) int {
	occurrences := make(map[string]int)
	
	for _, polymer := range polymerMap {
		for _, ch := range polymer {
			occurrences[string(ch)]++
		}
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

// func countElements(polymer string) int {
// 	occurrences := make(map[string]int)
	
// 	for _, ch := range polymer {
// 		occurrences[string(ch)]++
// 	}

// 	fmt.Println(occurrences)

// 	var mostCommon, leastCommon int

// 	for _, v := range occurrences {
// 		if leastCommon == 0 {
// 			leastCommon = v
// 		}

// 		if v > mostCommon {
// 			mostCommon = v
// 		} else if v < leastCommon {
// 			leastCommon = v
// 		}
// 	}

// 	return mostCommon - leastCommon
// }

func insertion(templateMap map[int]string, instructions map[int][]string) map[int]string {
	separated := make(map[int]string)
	finalMap := make(map[int]string)
	
	var newString string
	
	// Separate letters
	for _, template := range templateMap {
		for i, ch := range template {
			if i+1 < len(template) {
				newString = string(ch) + string(template[i+1])
			}
			
			separated[i] = newString
			newString = ""
		}
	}

	// Pair insertion. [N (N] [C) B] -> [N C (N] B [C) H B]
	counter := 0
	for i, str := range separated {
		for _, instruction := range instructions {
			if str == instruction[0] {
				myStr := string(str[0]) + instruction[1]
				if i == len(separated)-2 {
					myStr += string(str[1])
				}

				finalMap[counter] = myStr
				counter++

				// finalString = append(finalString, myStr)
			}
		}
	}

	for i := 0; i < len(finalMap); i++ {
		fmt.Println(finalMap[i])
	}

	return finalMap
}

// func insertion(template string, instructions map[int][]string) string {
// 	separated := make([]string, 0)
// 	finalString := make([]string, 0)
	
// 	var newString string
	
// 	// Separate letters
// 	for i, ch := range template {
// 		if i+1 < len(template) {
// 			newString = string(ch) + string(template[i+1])
// 		}
		
// 		separated = append(separated, newString)
// 		newString = ""
// 	}

// 	// Pair insertion. [N (N] [C) B] -> [N C (N] B [C) H B]
// 	for i, str := range separated {
// 		for _, instruction := range instructions {
// 			if str == instruction[0] {
// 				myStr := string(str[0]) + instruction[1]
// 				if i == len(separated)-2 {
// 					myStr += string(str[1])
// 				}
// 				finalString = append(finalString, myStr)
// 			}
// 		}
// 	}

// 	return strings.Join(finalString, "") 
// }

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

