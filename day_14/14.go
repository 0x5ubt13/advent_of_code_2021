package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	partOne()

	// Efficient part 2 to avoid running out of memory
	polymer, rules := getPuzzleInput("./14.in")

	// 1. Create the rules
	instructions := make(map[string]Output)
	for _, rule := range rules {
		var k, v string
		fmt.Sscanf(rule, "%s -> %s", &k, &v)

		instructions[k] = Output{
			Pair1: 		string(k[0]) + v,
			Pair2: 		v + string(k[1]),
			NewElement: v,
		}
	}

	// 2. Convert polymer to struct format
	s := newStep(polymer)

	// 3. Loop over 40 times
	for i := 0; i < 40; i++ {
		s = nextStep(s, instructions)
	}

	// 4. Solve puzzle
	var mostCommon, leastCommon int

	for _, v := range s.Count {
		if leastCommon == 0 {
			leastCommon = v
		}

		if v > mostCommon {
			mostCommon = v
		} else if v < leastCommon {
			leastCommon = v
		}
	}

	fmt.Printf("Part 2 -> %d\n", mostCommon - leastCommon)
}


type Step struct {
	Pairs map[string]int
	Count map[string]int
}

type Output struct {
	Pair1 string
	Pair2 string
	NewElement string
}

func nextStep(in *Step, rules map[string]Output) *Step {
	// New struct to hold the next step
	out := &Step{
		Pairs: make(map[string]int),
		Count: make(map[string]int),
	}

	// Copying the existing counts
	for k, v := range in.Count {
		out.Count[k] = v
	}

	// Output rules counting the insertions, expanding in different pairs
	for p, count := range in.Pairs {
		// For each pair: increment count of new pairs with insertions
		t := rules[p]
		out.Pairs[t.Pair1] += count
		out.Pairs[t.Pair2] += count
		out.Count[t.NewElement] += count
	}

	return out
}

func newStep(in string) *Step {
	step := &Step{
		Pairs: make(map[string]int), // pairs in the polymer and how many of each
		Count: make(map[string]int), // elements in the polymer and how many of each
	}

	// Take every pair in the polymer and add it to Pairs and Counts
	for i := 0; i < len(in)-1; i++ {
		pair := in[i : i+2]
		step.Pairs[pair]++
	}


	for i := 0; i < len(in); i++ {
		step.Count[string(in[i])]++
	}

	return step
}

func partOne() {
	polymer, rules := getPuzzleInput("./14.in")

	instructions := make(map[int][]string, 0)
	for i, rule := range rules {
		instructions[i] = strings.Split(rule, " -> ")
	}

	// Pair insertion. [N (N] [C) B] -> [N C (N] B [C) H B]
	for steps := 0; steps < 10; steps++ {
		polymer = insertion(polymer, instructions)
		fmt.Printf("After step %d:\t%d\n", steps, len(polymer))
		if steps == 9 {
			result := countElements(polymer)
			fmt.Printf("Part 1 -> %d\n", result)
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

