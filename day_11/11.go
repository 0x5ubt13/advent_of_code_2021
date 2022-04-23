package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	grid, flashed := getPuzzleInput("./11.in")
	part1 := 250
	flashes := 0

	maxY := len(grid)-1
	maxX := len(grid[0])-1
	
	loop:
	for steps := 0; steps < part1; steps++ {
		// Uncomment following snippet to graphically view the grid:
		// fmt.Println("Step", steps)
		// for i := 0; i < len(grid); i++ {
		// 	fmt.Println(grid[i])
		// }

		// 1: new step, add 1 to every octopus
		part2 := 0
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] == 0 {
					part2++
				}

				if part2 == 100 {
					fmt.Println("Part 2 ->", steps)
					break loop
				}
				// fmt.Printf("Row %d, Column %d, Value of the octopus -> %d\n", y, x, data[y][x])
				grid[y][x]++
			}
		}

		// 2: if higher than 9, make all neighbours +1, only once for step
		for z := 0; z < 20; z++ {
			// Make sure every single one of them has the chance to flash
			for y := 0; y < len(grid); y++ {
				for x := 0; x < len(grid[0]); x++ {
					
					positions := []string{"north", "northeast", "east", "southeast", "south", "southwest", "west", "northwest"}
					removing := make([]string, 0)

					if grid[y][x] > 9 && flashed[y][x] == false {
						// Add to flashed:
						flashed[y][x] = true
						flashes++

						// 4-way check:
						// Top row: eliminate north
						if y == 0 {
							removing = append(removing, "north", "northeast", "northwest")
						}

						// Bottom row: eliminate south
						if y == maxY {
							removing = append(removing, "south", "southeast", "southwest")
						}

						// First column: eliminate west
						if x == 0 {
							removing = append(removing, "northwest", "west", "southwest")
						}

						// Last column: eliminate east
						if x == maxX {
							removing = append(removing, "northeast", "east", "southeast")
						}

						positions = removeStrFromSlice(positions, removing)
						// fmt.Printf("Row %d, Column %d, Value of the octopus -> %d\n", y, x, grid[y][x])
						// fmt.Printf("\nPositions: %+v, Removing: %+v\n", positions, removing)
						for _, pos := range positions {
							switch pos {
							case "north":
								grid[y-1][x]++
							case "northeast":
								grid[y-1][x+1]++
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
		}

		// 3: if > 9, restore to 0 
		// (An octopus can only flash at most once per step.)
		
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				// fmt.Printf("Row %d, Column %d, Value of the octopus -> %d\n", y, x, grid[y][x])
				if flashed[y][x] == true {
					flashed[y][x] = false
					grid[y][x] = 0
				}
			}
		}
		
		if steps == 99 {
			fmt.Println("Part 1 ->", flashes)
		}
	}
}


func removeStrFromSlice(s []string, r []string) []string {
	for _, w := range r {
		for k, v := range s {
			if w == v {
				s = append(s[:k], s[k+1:]...)
			}
		}
	}
	
	return s
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