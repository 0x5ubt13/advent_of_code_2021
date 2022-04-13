package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	// signals[0] -> same line than digits[0]
	_, _, signals, digits := getPuzzleInput("./day_8_input.txt")

	decoded := make([]string, 0, 0)

	for line := 0; line < len(signals); line++ {
		// line by line
		var match int
		decoding := make(map[int]string)


		fmt.Printf("Line %d --> %v | %v\n", line, signals[line], digits[line])

		currentLine := strings.Split(strings.TrimSuffix(signals[line], "\r"), " ")

		for _, signal := range currentLine {
			// Make sure we populate 1, 4, 7 and 8 first
			switch len(signal) {
			case 2: //1
				decoding[1] = signal
			case 3: // 7
				decoding[7] = signal
			case 4: // 4
				decoding[4] = signal
			case 7: // 8
				decoding[8] = signal
			}
		}

		for _, signal := range currentLine {
			// Populate 0, 2, 3, 5, 6 and 9
			switch len(signal) {
			case 6: // 9, 0, 6
				// 9
				match = 0 
				for _, r := range signal {
					if rInStr(r, decoding[4]) == true {
						match++
					}
				}

				if match == 4 {
					decoding[9] = signal
					break
				}

				// 0
				match = 0
				for _, r := range signal {
					if rInStr(r, decoding[1]) == true {
						match++
					} 
				}

				if match == 2 {
					decoding[0] = signal
					break
				}

				// 6
				decoding[6] = signal
			
			case 5:
				// 3
				match = 0
				for _, r := range signal {
					if rInStr(r, decoding[1]) == true {
						match++
					}
				}

				if match == 2 {
					decoding[3] = signal
					break
				}
				
				// 5
				var match4, match7 int
				for _, r := range signal {
					if rInStr(r, decoding[7]) == true {
						match7++
					}

					if rInStr(r, decoding[4]) == true {
						match4++
					}
				}

				if match4 == 3 && match7 == 2 {
					decoding[5] = signal
					break
				}

				// 2
				decoding[2] = signal
			}
		}

		// decoding populated, now we need to decode the digits
		currentDigitsLine := strings.Split(strings.TrimSuffix(digits[line], "\r"), " ")
		digitsLine := ""
		for _, digit := range currentDigitsLine {
			// Make sure we populate 1, 4, 7 and 8 first
			switch len(digit) {
			// Easy ones
			case 2: //1
				digitsLine += "1"
			case 3: // 7
				digitsLine += "7"
			case 4: // 4
				digitsLine += "4"
			case 7: // 8
				digitsLine += "8"

			// Complicated ones
			case 5:
				var match3, match5, match2 int
				for _, r := range digit {
					for _, ru := range decoding[3] {
						if r == ru {
							match3++
						}
					}

					for _, ru := range decoding[5] {
						if r == ru {
							match5++
						}
					}

					for _, ru := range decoding[2] {
						if r == ru {
							match2++
						}
					}
				}

				if match2 == 5 {
					digitsLine += "2"
				} else if match3 == 5 {
					digitsLine += "3"
				} else if match5 == 5 {
					digitsLine += "5"
				} else {
					panic("You're doing something wrong")
				}
			case 6:
				var match0, match6, match9 int
				for _, r := range digit {
					for _, ru := range decoding[0] {
						if r == ru {
							match0++
						}
					}

					for _, ru := range decoding[6] {
						if r == ru {
							match6++
						}
					}

					for _, ru := range decoding[9] {
						if r == ru {
							match9++
						}
					}
				}

				if match0 == 6 {
					digitsLine += "0"
				} else if match6 == 6 {
					digitsLine += "6"
				} else if match9 == 6 {
					digitsLine += "9"
				} else {
					panic("You're doing something wrong")
				}
			}
		}

		decoded = append(decoded, digitsLine)

		for i := 0; i < len(currentLine); i++ {
			fmt.Printf("%s -> %d\n", decoding[i], i)
		}

		// End of line
	}

	// End of input
	
	fmt.Println(decoded)

	var solution int
	
	for _, number := range decoded {
		sol, err := strconv.Atoi(number)
		chkErr(err)

		solution += sol
	}

	fmt.Println(solution)
}

func findSignals(digit, signal string) int {
	matches := 0
	
	for _, r := range digit {
		for _, ru := range signal {
			if r == ru {
				matches++
			}
		}
	}

	return matches
}


func rInStr(x rune, y string) bool {
	for _, s := range y {
		if s == x {
			return true
		}
	}
	return false
}

func partOne() {
	_, digits, _, _ := getPuzzleInput("./day_8_input.txt")
	
	var ones, fours, sevens, eights, hits, part1 int

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

	part1 = ones + fours + sevens + eights

	fmt.Println("Part 1 ->", part1)
}

func getPuzzleInput(filename string) ([]string, []string, []string, []string) {
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

	return signals, digits, sig, dig
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}