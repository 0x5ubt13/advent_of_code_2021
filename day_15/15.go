package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	grid := getGrid("./15_test.txt")

	// for i := 0; i < len(grid); i++ {
	// 	fmt.Println(grid[i])
	// }

	maxY := len(grid)
	maxX := len(grid[0])
	var previous string
	var risk int

	queue := []Coord{{Y: 0, X: 0}}

	// for len(queue) > 0 {
	for i := 0; i < 10; i++ {
		cur := queue[0]
		queue = queue[1:]

		// Stop the queue if we reached maxY and maxX
		if cur.Y == maxY-1 && cur.X == maxX-1 {
			break
		}  

		// 4-way check:
		// First check if they exist before comparing
		existing := make(map[string]int, 0)

		// Top exists
		if cur.Y > 0 && previous != "top" {
			existing["top"] = grid[cur.Y - 1][cur.X] 
		}

		// Bottom exists
		if cur.Y < maxY-1 && previous != "bottom" {
			existing["bottom"] = grid[cur.Y + 1][cur.X]
		}

		// Left exists
		if cur.X > 0 && previous != "left" {
			existing["left"] = grid[cur.Y][cur.X - 1]
		} 

		// Right exists
		if cur.X < maxX-1 && previous != "right" {
			existing["right"] = grid[cur.Y][cur.X + 1]
		}

		// Compare the ones that exist
		minValue := 10
		safestPosition := ""

		for positionKey, riskValue := range existing {
			if riskValue < minValue {
				safestPosition = positionKey
			}
		}

		// Transform safest position to coordinates
		// and continue the path via the safest position
		nextCoord := Coord{}
		if safestPosition == "top" {

			nextCoord = Coord{Y: cur.Y-1, X: cur.X}
			previous = "bottom"
			risk += grid[nextCoord.Y][nextCoord.X]

			queue = append(queue, Coord{Y: cur.Y-1, X: cur.X})
			fmt.Println("new position:", nextCoord)

		} else if safestPosition == "bottom" {

			nextCoord = Coord{Y: cur.Y+1, X: cur.X}
			previous = "top"
			risk += grid[nextCoord.Y][nextCoord.X]

			queue = append(queue, nextCoord)
			fmt.Println("new position:", nextCoord)

		} else if safestPosition == "left" {

			nextCoord = Coord{Y: cur.Y, X: cur.X-1}
			previous = "right"
			risk += grid[nextCoord.Y][nextCoord.X]

			queue = append(queue, nextCoord)
			fmt.Println("new position:", nextCoord)
			
		} else if safestPosition == "right" {
			
			nextCoord = Coord{Y: cur.Y, X: cur.X+1}
			previous = "left"
			risk += grid[nextCoord.Y][nextCoord.X]
			queue = append(queue, nextCoord)

			fmt.Println("new position:", nextCoord)
		}

		
	
	}
	
	fmt.Println(risk)
}

type Coord struct {
	Y int
	X int
}

func getGrid(filename string) map[int][]int {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	rows := make(map[int][]int)
	grid := strings.Split(string(bytes), "\n")

	for i, line := range grid {
		for _, ch := range strings.TrimSuffix(line, "\r") {
			newInt, err := strconv.Atoi(string(ch))
			if err != nil {
				fmt.Println(err)
			}

			rows[i] = append(rows[i], newInt)

		}
	}

	return rows
}