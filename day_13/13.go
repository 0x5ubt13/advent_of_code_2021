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
	
	// Make use of the coordinates and draw marks
	for _, coord := range coords {
		grid[coord.Y][coord.X] = "#"
	}

	printGrid(grid)
	
	
	foldGrid(grid, folds[0], maxY, maxX)




}

func foldGrid(grid map[int][]string, f Fold, maxY, maxX int) {
	grid2 := make(map[int][]string)
	for k, v := range grid {
		grid2[k] = v
	}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if f.Axis == "y" && y == f.Position {
				grid2[y][x] = "-"
			}
		}
	}

	printGrid(grid2)
	
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
					} else if int(ch) > 40 {
						fmt.Println(ch)
						newFold.Position = int(ch) - 48
					}
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

