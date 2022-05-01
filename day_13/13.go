package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main () {
	coords, folds, maxX, maxY := getPuzzleInput("./13.test")

	for i, c := range coords {
		fmt.Println(i, c)
	}

	for i, f := range folds {
		fmt.Println(i, f)
	}

	// Create grid and populate with empty dots
	grid := make(map[int][]string)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			grid[y] = append(grid[y], ".")
		}
	}

	fmt.Printf("Max X -> %d\nMax Y -> %d\n", maxX, maxY)
	
	// Make use of the coordinates and draw marks
	for _, coord := range coords {
		grid[coord.Y][coord.X] = "#"
	}

	printGrid(grid)

	countMarks(grid, maxX, maxY)

	grid, maxY, maxX = foldGrid(grid, folds[0], maxY, maxX)

	printGrid(grid)

	countMarks(grid, maxX, maxY)

	grid, maxY, maxX = foldGrid(grid, folds[1], maxY, maxX)
	countMarks(grid, maxX, maxY)



}

func countMarks(grid map[int][]string, maxX, maxY int) {
	marks := 0
	for i:=0; i<len(grid)+1; i++ {
		marks += strings.Count(strings.Join(grid[i], ""), "#")
	}

	fmt.Println(marks)
}

func foldGrid(grid map[int][]string, f Fold, maxY, maxX int) (map[int][]string, int, int) {
	grid2 := make(map[int][]string)
	for k, v := range grid {
		fmt.Printf("Copying grid[%d] -> %v\n", k, grid[k])
		grid2[k] = v
	}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// If arrives to the folding point
			if f.Axis == "y" && y == f.Position || f.Axis == "x" && x == f.Position {
				fmt.Printf("Y -> %d, X -> %d, maxY -> %d, maxX -> %d\n", y, x, maxY, maxX)

				grid2[y][x] = "-"
			} 
			
			// If goes beyond the folding point
			if f.Axis == "y" && f.Position < y && grid2[y][x] == "#" {
				diff := y - f.Position
				grid2[f.Position - diff][x] = "#"
			} 

			if f.Axis == "x" && f.Position < x && grid2[y][x] == "#" {
				diff := x - f.Position
				grid2[y][f.Position - diff] = "#"
			}


		}
	}

	if f.Axis == "y" {
		for y := 0; y < maxY; y++ {
			if y >= f.Position {
				delete(grid2, y)
			}
		}

		maxY = f.Position+1
	} else {
		for x := 0; x < maxX; x++ {
			if x >= f.Position {
				delete(grid2, x)
			}
		}
		maxX = f.Position+1
	}
	
	return grid2, maxY, maxX
	
}

func printGrid(grid map[int][]string) {
	fmt.Println()
	for i:=0; i<len(grid); i++{
		fmt.Println(grid[i])
	}
	fmt.Println()
}

func drawGrid(maxX, maxY int, coords []Coord) map[int][]string {
	grid := make(map[int][]string)
	
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			grid[y] = append(grid[y], ".")
		}
	}

	return grid
}

type Coord struct {
	X int
	Y int
}

type Fold struct {
	Axis string
	Position int
}

func getPuzzleInput(fn string) ([]Coord, []Fold, int, int) {
	bytes, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	coords := make([]Coord, 0)
	folds := make([]Fold, 0)

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		var newCoord Coord
		var newFold Fold
		var num = []string{}

		coordinates := strings.Split(line, ",")
		newCoord.X, err = strconv.Atoi(strings.TrimSuffix(coordinates[0], "\r"))
		if err != nil {
			fmt.Println(err)
			if len(line) > 1 {
				for _, ch := range line {
					if ch == 'y' {
						newFold.Axis = "y"
					} else if ch == 'x' {
						newFold.Axis = "x"
					} 
				}

				num = strings.Split((strings.TrimSuffix(line, "\r")), "=")
				newFold.Position, err = strconv.Atoi(num[1])
				if err != nil {
					fmt.Println(err)
				}
				
				folds = append(folds, newFold)
			}
			continue
		} 
		newCoord.Y, err = strconv.Atoi(strings.TrimSuffix(coordinates[1], "\r"))
		if err != nil {
			fmt.Println(err)
		}
		
		coords = append(coords, newCoord)
	}

	var maxX, maxY int

	for i, f := range folds {
		if i == 0 || i == 1 {
			if f.Axis == "y" {
				maxY = f.Position * 2 + 1
			} else {
				maxX = f.Position * 2 + 1
			}
		} else {
			break
		}
	}

	return coords, folds, maxX, maxY
}

