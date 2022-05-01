package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	template, rules := getPuzzleInput("./14.test")

	fmt.Println(template)
	for _, l := range rules {
		fmt.Printf("%v\n", l)
	}

	// Pair insertion. [N (N] [C) B] -> [N C (N] B [C) H B]
	result := insertion(template, rules)
	fmt.Println(result)
}

func insertion(template string, rules []string) string {
	// Check the template and grab elements 2 by 2
	for _, ch := range template {

	}
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

