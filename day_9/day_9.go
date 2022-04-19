package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	heightMap := getPuzzleInput("9.in")

	// partOne(heightMap)
	smarterPartOne(heightMap)
	PartTwo(heightMap)
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

func smarterPartOne(heightMap map[int][]int) {
	maxY := len(heightMap)
	maxX := len(heightMap[0])
	riskRating := 0

	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[y]); x++ {
			// fmt.Println("Row", y, "Column", x, "->", heightMap[y][x])
			pointer := heightMap[y][x]
			// 4-way check: top, bottom, left and right boundaries:

			// if not uppermost row and "up" is greater or equal than the pointer, pass
			if y > 0 && heightMap[y-1][x] <= pointer {
				continue
			}

			// if not bottommost row and "down" is greater or equal than the pointer, pass
			if y < maxY-1 && heightMap[y+1][x] <= pointer {
				continue
			}

			// if not leftmost column and "left" is greater or equal than the pointer, pass
			if x > 0 && heightMap[y][x-1] <= pointer {
				continue
			}

			// if not rightmost column and "right" is greater or equal than the pointer, pass
			if x < maxX-1 && heightMap[y][x+1] <= pointer {
				continue
			}

			// If made it through all conditions, score
			riskRating += pointer + 1
		}
	}

	fmt.Printf("Part 1 -> %d\n", riskRating)

}

func PartTwo(heightMap map[int][]int) {
	maxY := len(heightMap)
	maxX := len(heightMap[0])

	var grid2 [][]rune = make([][]rune, len(heightMap))	
	for i, row := range heightMap {
		grid2[i] = make([]rune, len(row))
	}

	name := 'A'

	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[y]); x++ {
			if grid2[y][x] != 0 {
				continue
			}

			if heightMap[y][x] == 9 {
				grid2[y][x] = '9'
				continue
			}

			type Point struct {
				x, y int
			}

			// Branch off to a queue, finding all coords of the basin (breadth-first search: https://en.wikipedia.org/wiki/Breadth-first_search)
			queue := []Point{{x, y}}

			// if queue isn't empty
			for len(queue) > 0 {
				// grab the first point, then pop it off the queue
				p := queue[0]
				queue = queue[1:]

				grid2[p.y][p.x] = name

				if p.y > 0 && heightMap[p.y - 1][p.x] != 9 && grid2[p.y - 1][p.x] == 0 {
					queue = append(queue, Point{p.x, p.y - 1})
				}
	
				if p.y < maxY - 1 && heightMap[p.y + 1][p.x] != 9 && grid2[p.y + 1][p.x] == 0 {
					queue = append(queue, Point{p.x, p.y + 1})
				}
	
				if p.x > 0 && heightMap[p.y][p.x - 1] != 9 && grid2[p.y][p.x - 1] == 0 {
					queue = append(queue, Point{p.x - 1, p.y})
				}
	
				if p.x < maxX - 1 && heightMap[p.y][p.x + 1] != 9 && grid2[p.y][p.x + 1] == 0 {
					queue = append(queue, Point{p.x + 1, p.y})
				}
			}
			
			// Queue over, change name of next basin
			name++
		}
	}

	basins := make(map[rune]int)

	for _, v := range grid2 {
		for _, r := range v {
			if r != '9' {
				basins[r]++
			}
		}
	}

	var sorting []int

	for _, v := range basins {
		sorting = append(sorting, v)
	}

	sort.Ints(sorting)
	fmt.Printf("Part 2 -> %d", sorting[len(sorting)-1] * sorting[len(sorting)-2] * sorting[len(sorting)-3])
}

func getPuzzleInput(filename string) (map[int][]int) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")

	intLines := make(map[int][]int, len(lines))
	for i, line := range lines {

		for _, num := range line {
			casting, err := strconv.Atoi(strings.TrimSuffix(string(num), "\r"))
			if err != nil {
				// fmt.Println(err)
			} else {
				intLines[i] = append(intLines[i], casting)
			}
		}

	}

	return intLines
}