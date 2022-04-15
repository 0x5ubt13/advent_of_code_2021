package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	heightMap, shadowMap := getPuzzleInput("9.test")

	partOne(heightMap)
	partTwo(heightMap, shadowMap)

}

func partOne(heightMap map[int][]int) {
	riskLevel := 0

	for row := 0; row < len(heightMap); row++ {
		// Outer loop 
		for column := 0; column < len(heightMap[0]); column++ {
			// Inner loop. Column hit. Print where are we
			pointer := heightMap[row][column]

			// Start row conditions:
			if row == 0 {
				// Top row. Condition -> no up

				// Start column conditions:
				if column == 0 {
					// First column. Condition -> no left

					down	:= heightMap[row+1][column]
					right	:= heightMap[row][column+1]

					if down > pointer && right > pointer {
						riskLevel += pointer + 1
					}

				} else if column == len(heightMap[0]) - 1 {
					// Last column. Condition -> no right

					down	:= heightMap[row+1][column]
					left	:= heightMap[row][column-1]

					if down > pointer && left > pointer {
						riskLevel += pointer + 1
					}

				} else {
					// Any other column in the middle

					down	:= heightMap[row+1][column]
					left	:= heightMap[row][column-1]
					right	:= heightMap[row][column+1]

					if down > pointer && left > pointer && right > pointer {
						riskLevel += pointer + 1
					}
				}

			} else if row == len(heightMap) - 1 { 
				// Bottom row. Condition -> no down

				// Start column conditions:
				if column == 0 {
					// First column. Condition -> no left

					up 		:= heightMap[row-1][column]
					right	:= heightMap[row][column+1]

					if up > pointer && right > pointer {
						riskLevel += pointer + 1
					}

				} else if column == len(heightMap[0]) - 1 {
					// Last column. Condition -> no right

					up 		:= heightMap[row-1][column]
					left	:= heightMap[row][column-1]

					if up > pointer && left > pointer {
						riskLevel += pointer + 1
					}

				} else {
					// Any other column in the middle

					up 		:= heightMap[row-1][column]
					left	:= heightMap[row][column-1]
					right	:= heightMap[row][column+1]

					if up > pointer && left > pointer && right > pointer {
						riskLevel += pointer + 1
					}
				}

			} else {
				// Any other row in the middle. Both up and down operative

				// Start column conditions:
				if column == 0 {
					// First column. Condition -> no left

					up 		:= heightMap[row-1][column]
					down	:= heightMap[row+1][column]
					right	:= heightMap[row][column+1]

					if up > pointer && down > pointer && right > pointer {
						riskLevel += pointer + 1
					}

				} else if column == len(heightMap[0]) - 1 {
					// Last column. Condition -> no right

					up 		:= heightMap[row-1][column]
					down	:= heightMap[row+1][column]
					left	:= heightMap[row][column-1]

					if up > pointer && down > pointer && left > pointer {
						riskLevel += pointer + 1
					}

				} else {
					// Any other column in the middle

					up 		:= heightMap[row-1][column]
					down	:= heightMap[row+1][column]
					left	:= heightMap[row][column-1]
					right	:= heightMap[row][column+1]

					if up > pointer && down > pointer && left > pointer && right > pointer {
						riskLevel += pointer + 1
					}
				}
			}
		}

		// end of row
	} 

	// end of matrix
	fmt.Printf("Part 1 -> %d", riskLevel)

}

func partTwo(heightMap map[int][]int, shadowMap map[int][]bool) {
	for row := 0; row < len(heightMap); row++ {
		// Outer loop 
		for column := 0; column < len(heightMap[0]); column++ {
			// Inner loop. Column hit. Print where are we
			// fmt.Println("Row", row, "Column", column, "->", heightMap[row][column], shadowMap[row][column])
			
			// intPointer := heightMap[row][column]
			// boolPointer := shadowMap[row][column]
			// basinPointer := heightmap[row][column]
			// basinCounter := 0

			// Start row conditions:
			if row == 0 {
				// Top row. Condition -> no up

				// Start column conditions:
				if column == 0 {
					// First column. Condition -> no left
				} else if column == len(heightMap[0]) - 1 {
					// Last column. Condition -> no right
				} else {
					// Any other column in the middle
				}

			} else if row == len(heightMap) - 1 { 
				// Bottom row. Condition -> no down

				// Start column conditions:
				if column == 0 {
					// First column. Condition -> no left
				} else if column == len(heightMap[0]) - 1 {
					// Last column. Condition -> no right
				} else {
					// Any other column in the middle

				}

			} else {
				// Any other row in the middle. Both up and down operative

				// Start column conditions:
				if column == 0 {
					// First column. Condition -> no left
				} else if column == len(heightMap[0]) - 1 {
					// Last column. Condition -> no right
				} else {
					// Any other column in the middle

					up 		:= heightMap[row-1][column]
					down	:= heightMap[row+1][column]
					left	:= heightMap[row][column-1]
					right	:= heightMap[row][column+1]

				}
			}

			// Create an indefinite loop
			// if any of the surrounding are lower than the pointer,

			// Idea -> all basins are surrounded by 9s

		}

		// end of row
	} 

	// end of matrix

}


func getPuzzleInput(filename string) (map[int][]int, map[int][]bool) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")

	boolLines := make(map[int][]bool, len(lines))
	intLines := make(map[int][]int, len(lines))
	for i, line := range lines {

		for _, num := range line {
			casting, err := strconv.Atoi(strings.TrimSuffix(string(num), "\r"))
			if err != nil {
				// fmt.Println(err)
			} else {
				intLines[i] = append(intLines[i], casting)
				boolLines[i] = append(boolLines[i], false)
			}
		}

	}

	return intLines, boolLines

}