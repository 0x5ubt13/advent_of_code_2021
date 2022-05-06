package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	Y, X int
}

func main() {
	grid := getGrid("./15_test.txt")
	maxY := len(grid)
	maxX := len(grid[0])

	// Dijksra's algorithm:
	// 1: Mark all nodes unvisited. 
	// 1.2. Create a set of all the unvisited nodes called the unvisited set.

	visited := make(map[Coord]bool)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			newCoord := Coord{Y: y, X: x}
			visited[newCoord] = false 
		}
	}

	// 2. Assign to every node a tentative distance value: set it to zero for our initial node and to infinity for all other nodes. 
	//    The tentative distance of a node v is the length of the shortest path discovered so far between the node v and the starting node. 
	//    Since initially no path is known to any other vertex than the source itself (which is a path of length zero), all other tentative distances are initially set to infinity. 
	//    Set the initial node as current.
	tentativeDistance := math.MaxInt
	current := Coord{0, 0}


	// 3. For the current node, consider all of its unvisited neighbors and calculate their tentative distances through the current node. 
	//    Compare the newly calculated tentative distance to the one currently assigned to the neighbor and assign it the smaller one. 
	//    For example:
	//       if the current node A is marked with a distance of 6, and the edge connecting it with a neighbor B has length 2, then the distance to B through A will be 6 + 2 = 8. 
	// 	  If B was previously marked with a distance greater than 8 then change it to 8. Otherwise, the current value will be kept.
	neighbours := checkNeighbours(current, maxY, maxX)

	// 4. When we are done considering all of the unvisited neighbors of the current node, mark the current node as visited and remove it from the unvisited set. 
	//    A visited node will never be checked again (this is valid and optimal in connection with the 
	//    behavior in step 6.: that the next nodes to visit will always be in the order of 'smallest distance from initial node first' 
	//    so any visits after would have a greater distance).
	
	// 5. If the destination node has been marked visited (when planning a route between two specific nodes) or if the smallest tentative distance among the nodes in the 
	//    unvisited set is infinity (when planning a complete traversal; occurs when there is no connection between the initial node and remaining unvisited nodes), then stop. 
	//    The algorithm has finished.
	
	// 6. Otherwise, select the unvisited node that is marked with the smallest tentative distance, set it as the new current node, and go back to step 3.

}

// Check the neighbours of the current coordinate with a 4-way check
func checkNeighbours(current Coord, maxY, maxX int) []Coord {
	neighbours := make([]Coord, 0)
	
	if current.Y > 0  { 		// Top exists
		neighbours = append(neighbours, Coord{current.Y-1, current.X}) 
	}

	if current.Y < maxY-1 {		// Bottom exists
		neighbours = append(neighbours, Coord{current.Y+1, current.X}) 
	}

	if current.X > 0 { 			// Left exists
		neighbours = append(neighbours, Coord{current.Y, current.X-1}) 
	} 

	if current.X < maxX-1 {		// Right exists
		neighbours = append(neighbours, Coord{current.Y, current.X+1}) 
	}
	
	return neighbours
}


// // Starting over
// func oldmain() {
// 	maxY := len(grid)
// 	maxX := len(grid[0])
// 	var previous string
// 	var risk int

// 	queue := []Coord{{Y: 0, X: 0}}

// 	// for len(queue) > 0 {
// 	for i := 0; i < 10; i++ {
// 		cur := queue[0]
// 		queue = queue[1:]

// 		// Stop the queue if we reached maxY and maxX
// 		if cur.Y == maxY-1 && cur.X == maxX-1 {
// 			break
// 		}  

// 		// 4-way check:
// 		// First check if they exist before comparing
// 		existing := make(map[string]int, 0)

// 		// Top exists
// 		if cur.Y > 0 && previous != "top" {
// 			existing["top"] = grid[cur.Y - 1][cur.X] 
// 		}

// 		// Bottom exists
// 		if cur.Y < maxY-1 && previous != "bottom" {
// 			existing["bottom"] = grid[cur.Y + 1][cur.X]
// 		}

// 		// Left exists
// 		if cur.X > 0 && previous != "left" {
// 			existing["left"] = grid[cur.Y][cur.X - 1]
// 		} 

// 		// Right exists
// 		if cur.X < maxX-1 && previous != "right" {
// 			existing["right"] = grid[cur.Y][cur.X + 1]
// 		}

// 		// Compare the ones that exist
// 		minValue := 10
// 		safestPosition := ""

// 		for positionKey, riskValue := range existing {
// 			if riskValue < minValue {
// 				safestPosition = positionKey
// 			}
// 		}

// 		// Transform safest position to coordinates
// 		// and continue the path via the safest position
// 		nextCoord := Coord{}
// 		if safestPosition == "top" {

// 			nextCoord = Coord{Y: cur.Y-1, X: cur.X}
// 			previous = "bottom"
// 			risk += grid[nextCoord.Y][nextCoord.X]

// 			queue = append(queue, Coord{Y: cur.Y-1, X: cur.X})
// 			fmt.Println("new position:", nextCoord)

// 		} else if safestPosition == "bottom" {

// 			nextCoord = Coord{Y: cur.Y+1, X: cur.X}
// 			previous = "top"
// 			risk += grid[nextCoord.Y][nextCoord.X]

// 			queue = append(queue, nextCoord)
// 			fmt.Println("new position:", nextCoord)

// 		} else if safestPosition == "left" {

// 			nextCoord = Coord{Y: cur.Y, X: cur.X-1}
// 			previous = "right"
// 			risk += grid[nextCoord.Y][nextCoord.X]

// 			queue = append(queue, nextCoord)
// 			fmt.Println("new position:", nextCoord)
			
// 		} else if safestPosition == "right" {
			
// 			nextCoord = Coord{Y: cur.Y, X: cur.X+1}
// 			previous = "left"
// 			risk += grid[nextCoord.Y][nextCoord.X]
// 			queue = append(queue, nextCoord)

// 			fmt.Println("new position:", nextCoord)
// 		}

// 	}
	
// 	fmt.Println(risk)
// }


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