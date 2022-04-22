package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	grid, _ := getPuzzleInput("./11.test")
	part1 := 3

	maxY := len(grid)
	maxX := len(grid[0])
		
	for steps := 0; steps < part1; steps++ {
		fmt.Println("Step", steps)
		for i := 0; i < len(grid); i++ {
			fmt.Println(grid[i])
		}

		// 1: new step, add 1 to every octopus
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				// fmt.Printf("Row %d, Column %d, Value of the octopus -> %d\n", y, x, data[y][x])
				grid[y][x]++
			}
		}

		// 2: if 9 or higher, make all neighbours +1
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				positions := []string{"north", "northeast", "east", "southeast", "south", "southwest", "west", "northwest"}
				// positions := []string{""}
				if grid[y][x] > 9 {
					// top-left corner: Only east, southeast, south
					if y == 0 {
						positions = append(positions, "south")
					}

					if x == maxX {

					}



					for _, pos := range positions {
						switch pos {
						case "north":
							grid[y+1][x]++
						case "northeast":
							grid[y+1][x+1]++
						case "east":
							grid[y][x+1]++
						case "southeast":
							grid[y+1][x+1]++
						case "south":
							grid[y+1][x]++
						case "southwest":
							grid[y+1][x-1]++
						case "west":
							grid[y][x-1]++
						case "northwest":
							grid[y-1][x-1]++
						}
					}
					
					

				}
			}
		}

		// 2: any octopus with an energy level greater than 9 flashes. 
		// This increases the energy level of all adjacent octopuses by 1, including octopuses that are diagonally adjacent. 
		// If this causes an octopus to have an energy level greater than 9, it also flashes. 
		// This process continues as long as new octopuses keep having their energy level increased beyond 9. 
		// (An octopus can only flash at most once per step.)
		
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				// fmt.Printf("Row %d, Column %d, Value of the octopus -> %d\n", y, x, grid[y][x])
			}
		}
	}
		
	
	
}

func removeStrFromSlice(s []string, r []string) []string {
	for _, w := range r {
		for k, v := range s {
			if w == v {
				return append(s[:k], s[k+1:])
			}
		}
	}
}


func getPuzzleInput(filename string) (map[int][]int, map[int][]bool){
	bytes, _ := ioutil.ReadFile(filename)
	bools := make(map[int][]bool)
	ints := make(map[int][]int, 0)

	for i, line := range strings.Split(string(bytes), "\n") {
		for _, digit := range line {
			id, err := strconv.Atoi(string(digit))
			if err != nil {
			} else {
				ints[i] = append(ints[i], id)
				bools[i] = append(bools[i], false)
			}
		}
	}

	return ints, bools
}