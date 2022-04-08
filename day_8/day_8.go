package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	_, digits := getPuzzleInput("./day_8_input.txt")
	
	var ones, fours, sevens, eights int
	hits := 0

	for _, digit := range digits {
		if len(digit) == 3 || len(digit) == 2 || len(digit) == 4 || len(digit) == 7 {
			hits += 1
		}

		switch len(digit) {
		case 2: // 1
			ones++
		case 3: // 7
			sevens++
		case 4: // 4
			fours++
		case 7: // 8
			eights++
		}
	}

	part1 := ones + fours + sevens + eights

	fmt.Println("Part 1 ->", part1)
}

func getPuzzleInput(filename string) ([]string, []string) {
	bytes, err := ioutil.ReadFile(filename)
	chkErr(err)

	lines := strings.Split(string(bytes), "\n")
	
	var sig, signals, dig, digits []string

	for _, line := range lines {
		sig = append(sig, strings.Split(line, " | ")[0])
		dig = append(dig, strings.Split(line, " | ")[1])
	}

	for _, signal := range sig {
		signals = append(signals, strings.Split(strings.TrimSuffix(signal, "\r"), " ")...)
	} 

	for _, digit := range dig {
		digits = append(digits, strings.Split(strings.TrimSuffix(digit, "\r"), " ")...)
	} 

	return signals, digits
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}