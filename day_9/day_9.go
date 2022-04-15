package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	heightMap, shadowMap := getPuzzleInput("9.test")

	partOne(heightMap, shadowMap)

}

func partOne(heightMap map[int][]int, shadowMap map[int][]bool) {
	for row := 0; row < len(heightMap); row++ {
		// Outer loop 
		for column := 0; column < len(heightMap[0]); column++ {
			// Inner loop. Column hit. Print where are we
			fmt.Println("Row", row, "Column", column, "->", heightMap[row][column], shadowMap[row][column])

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
				// Any other row in the middle

				// Start column conditions:
				if column == 0 {
					// First column. Condition -> no left
				} else if column == len(heightMap[0]) - 1 {
					// Last column. Condition -> no right
				} else {
					// Any other column in the middle
				}
			}
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