package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	diagram := drawDiagram()

	instructions := getPuzzleInput("./day_5_input.txt") 

	var x1, x2, y1, y2 int
	var part2 bool

	for {
		for i, instruct := range instructions {
			if i % 2 == 0 {
				// Every even index is "x1, y1"
				x1, _ = strconv.Atoi(string(instruct[0]))
				y1, _ = strconv.Atoi(string(instruct[1]))
				// fmt.Printf("x1 = %v, y1 = %v, ", x1, y1)
			} else if i % 2 == 1 {
				// Every odd index is "x2, y2"
				x2, _ = strconv.Atoi(string(instruct[0]))
				y2, _ = strconv.Atoi(string(instruct[1]))
				// fmt.Printf("x2 = %v, y2 = %v\n", x2, y2)
			}

			if i % 2 == 1 {
				// Trigger only every second iteration to make sure we have updated x and y values

				// X axis
				if x1 == x2 {
					if y1 > y2 {
						// Rightwards
						// fmt.Println("Rightwards")
						for y := y2; y <= y1; y++ {
							// fmt.Println("x", x2, "y", y1)
							diag := &diagram[x1][y]
							*diag += 1
						}
						// diagram.print()
					} else if y1 < y2 {
						// Leftwards
						// fmt.Println("Leftwards")
						for y := y2; y >= y1; y-- {
							// fmt.Println("x", x2, "y", y1)
							diag := &diagram[x1][y]
							*diag += 1
						}
						// diagram.print()
					}
				}
				// Y axis
				if y1 == y2 {
					if x1 > x2 {
						// Downwards
						// fmt.Println("Downwards")
						for x2 <= x1 {
							// fmt.Println("x", x2, "y", y1)
							diag := &diagram[x2][y1]
							*diag += 1

							x2++
						}
						// diagram.print()
					} else if x1 < x2 {
						// Upwards
						// fmt.Println("Upwards")
						for x2 >= x1 {
							// fmt.Println("x", x2, "y", y1)
							diag := &diagram[x2][y1]
							*diag += 1

							x2--
						}
						// diagram.print()
					}
				}

				if x1 != x2 && y1 != y2 && part2 == true {
					// Diagonal
					fmt.Println("Diagonal: ")
					fmt.Printf("x1 = %v, y1 = %v, x2 = %v, y2 = %v\n", x1, y1, x2, y2)

					if x1 < x2 && y1 < y2 {
						for x1 <= x2 {
							diag := &diagram[x1][y1]
							*diag += 1
							
							x1++
							y1++
						}
					} else if x1 > x2 && y1 > y2 {
						for x1 >= x2 {
							diag := &diagram[x1][y1]
							*diag += 1

							x1--
							y1--
						}
					} else if x1 < x2 && y1 > y2 {
						for x1 <= x2 {
							diag := &diagram[x1][y1]
							*diag += 1

							x1++
							y1--
						}
					} else if x1 > x2 && y1 < y2 {
						for x1 >= x2 {
							diag := &diagram[x1][y1]
							*diag += 1

							x1--
							y1++
						}
					}
				}
			}
		}

		total := 0
		for i := 0; i < 1000; i++ {
			for j := 0; j < 1000; j++ {
				if diagram[i][j] >= 2 {
					total += 1
				}
			}
		}

		if part2 == false {
			fmt.Println("Part 1 ->", total)
			part2 = true
			total = 0
			diagram = drawDiagram()
			continue
		} else {
			fmt.Println("Part 2 ->", total)
			break
		}
	}
}

type Diagram [1000][1000]int

func (d Diagram) print() {
	for index, row := range d {
		fmt.Println(index, row)
	}
}

func drawDiagram() Diagram {
	var diagram Diagram

	return diagram
}

func getPuzzleInput(filename string) [][]string {
	bytes, err := ioutil.ReadFile(filename)
	chkErr(err)

	lines := strings.Split(strings.ReplaceAll(string(bytes), " -> ", " "), "\n")
	instructions := make([]string, 0)

	for _, line := range lines {
		instructions = append(instructions, strings.Split(line, " ")...)
	}
	
	var ins [][]string

	for _, line := range lines {
		insts := strings.Split(line, " ")
		for _, instruc := range insts {
			ins = append(ins, strings.Split(strings.TrimSuffix(instruc, "\r"), ","))
		}
	}	

	return ins
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
