package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	signals, digits := getPuzzleInput("./day_8_input.txt")
	
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

	// Part 2
	decoding := make(map[string]int, len(signals)
	decoded := make(map[int]int, len(signals))

	// signals[0] pairs with digits[0]
	for i:= 0; i < len(signals); i++ {

		for _, signal := range signals {
			switch len(signal) {
			case 2: // 1
				decoding[signal] = 1
			case 3: // 7
				decoding[signal] = 7
			case 4: // 4
				decoding[signal] = 4
			case 5: // 2, 3, 5
				for k, v := range decoding { 
					match := 0
					if v == 1 { // 3
						for _, letter := range k {
							for _, lett := range signal {
								if letter == lett {
									match += 1
								}
							}
						}

						if match == 2 {
							decoding[signal] = 3
						}

					} else if v == 4 || v == 7 {
						
					}
				}

			case 6: // 0, 6, 9 

			case 7: // 8
				decoding[signal] = 8
			case 
			}
		}

		
	}
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