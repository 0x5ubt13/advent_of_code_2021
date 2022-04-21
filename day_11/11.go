package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data := getPuzzleInput("./11.test")

	step := 0
	
	// First step, add 1 to every octopus
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {
			// fmt.Printf("Row %d, Column %d, Value of the octopus -> %d\n", y, x, data[y][x])
			data[y][x]++
		}
	}

	// Second step, any octopus with an energy level greater than 9 flashes. 
	// This increases the energy level of all adjacent octopuses by 1, including octopuses that are diagonally adjacent. 
	// If this causes an octopus to have an energy level greater than 9, it also flashes. 
	// This process continues as long as new octopuses keep having their energy level increased beyond 9. 
	// (An octopus can only flash at most once per step.)
	// for y := 0; y < len(data); y++ {
	// 	for x := 0; x < len(data[0]); x++ {
	// 		fmt.Printf("Row %d, Column %d, Value of the octopus -> %d\n", y, x, data[y][x])
	// 	}
	// }


}

func getPuzzleInput(filename string) map[int][]int {
	bytes, _ := ioutil.ReadFile(filename)

	ints := make(map[int][]int, 0)
	for i, line := range strings.Split(string(bytes), "\n") {
		for _, digit := range line {
			id, _ := strconv.Atoi(string(digit))
			ints[i] = append(ints[i], id)
		}
	}

	return ints
}