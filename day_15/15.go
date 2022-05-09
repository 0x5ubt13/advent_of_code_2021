package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type Coord struct {
	Y, X int
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	grid := getGrid("./15_in.txt")
	maxY := len(grid)
	maxX := len(grid[0])

	part1 := time.Now()
	partOne(grid, maxY, maxX)
	elapsed := time.Since(part1)
	log.Printf("%s took %s", "Part1", elapsed)


	
	maxY2 := maxY * 5
	maxX2 := maxX * 5
	defer timeTrack(time.Now(), "part2")
	partTwo(grid, maxY2, maxX2)

}

func partTwo(grid map[int][]int, maxY, maxX int) {
	nodeDistance := make(map[Coord]int)
	visited := make(map[Coord]bool)

	tentativeDistance := 999999999999
	
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// "Zero for out initial node"
			if y == 0 && x == 0 {
				nodeDistance[Coord{y, x}] = 0
			} else {
				nodeDistance[Coord{y, x}] = tentativeDistance
			}
		}
	}

	rows := len(grid)
	cols := len(grid[0])
	
	current := Coord{0,0}

	for {
		if visited[current] == true {
			continue
		}

		neighbours := checkNeighbours(current, maxY, maxX)
		for _, neighbour := range neighbours {
			if visited[neighbour] == true {
				continue
			}

			// Part 2 special: increment the grid 5 times in both dimensions using modulus

			y := neighbour.Y % rows // capping to length of rows
			x := neighbour.X % cols // capping to length of cols
			val := grid[y][x] + neighbour.Y / rows + neighbour.X / cols // update values accordingly
			if val > 9 { 
				val -= 9 
			} 

			newDistance := nodeDistance[current] + val
			if newDistance < nodeDistance[neighbour] {
				nodeDistance[neighbour] = newDistance
			}
		}

		visited[current] = true
		
		if visited[Coord{maxY-1, maxX-1}] == true {
			fmt.Println(nodeDistance[Coord{maxY-1, maxX-1}])
			break
		} else {
			minCoord := Coord{maxY, maxX}
			minDistance := math.MaxInt

			for k, v := range nodeDistance {
				// fmt.Println(k, v)
				if v < minDistance && !visited[k] {
					minDistance = v
					minCoord = k
				}
			}

			current = minCoord
		}
	}
}

func partOne(grid map[int][]int, maxY, maxX int) {
	// Dijksra's algorithm:
	// 1: Mark all nodes unvisited.
	// 1.2. Create a set of all the unvisited nodes called the unvisited set.
	nodeDistance := make(map[Coord]int)

	// 2. Assign to every node a tentative distance value: set it to zero for our initial node and to infinity for all other nodes.
	//    The tentative distance of a node v is the length of the shortest path discovered so far between the node v and the starting node.
	//    Since initially no path is known to any other vertex than the source itself (which is a path of length zero),
	//    all other tentative distances are initially set to infinity.
	//    Set the initial node as current.
	tentativeDistance := math.MaxInt

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// "Zero for out initial node"
			if y == 0 && x == 0 {
				nodeDistance[Coord{y, x}] = 0
			} else {
				nodeDistance[Coord{y, x}] = tentativeDistance
			}
		}
	}
	visited := make(map[Coord]bool)

	current := Coord{0, 0}

	for {
		// fmt.Println(current)

		if visited[current] == true {
			continue
		}

		// 3. For the current node, consider all of its unvisited neighbors and calculate their tentative distances through the current node.
		//    Compare the newly calculated tentative distance to the one currently assigned to the neighbor and assign it the smaller one.
		//    For example:
		//       if the current node A is marked with a distance of 6, and the edge connecting it with a neighbor B has length 2, then the distance to B through A will be 6 + 2 = 8.
		// 	  If B was previously marked with a distance greater than 8 then change it to 8. Otherwise, the current value will be kept.
		neighbours := checkNeighbours(current, maxY, maxX)
		for _, neighbour := range neighbours {
			if visited[neighbour] == true {
				continue
			}

			newDistance := nodeDistance[current] + grid[neighbour.Y][neighbour.X]

			if newDistance < nodeDistance[neighbour] {
				nodeDistance[neighbour] = newDistance
			}

		}

		// 4. When we are done considering all of the unvisited neighbors of the current node, mark the current node as visited and remove it from the unvisited set.
		//    A visited node will never be checked again (this is valid and optimal in connection with the
		//    behavior in step 6.: that the next nodes to visit will always be in the order of 'smallest distance from initial node first'
		//    so any visits after would have a greater distance).
		visited[current] = true
		
		// 5. If the destination node has been marked visited (when planning a route between two specific nodes) or if the smallest tentative distance among the nodes in the
		//    unvisited set is infinity (when planning a complete traversal; occurs when there is no connection between the initial node and remaining unvisited nodes), then stop.
		//    The algorithm has finished.
		if visited[Coord{maxY-1, maxX-1}] == true {
			fmt.Println(nodeDistance[Coord{maxY-1, maxX-1}])
			break
		} else {
			// 6. Otherwise, select the unvisited node that is marked with the smallest tentative distance, set it as the new current node, and go back to step 3.
			minCoord := Coord{maxY, maxX}
			minDistance := math.MaxInt

			for k, v := range nodeDistance {
				// fmt.Println(k, v)
				if v < minDistance && !visited[k] {
					minDistance = v
					minCoord = k
				}
			}

			current = minCoord
			// fmt.Println("new current", current)
		}
	}
}

// Check the neighbours of the current coordinate with a 4-way check
func checkNeighbours(current Coord, maxY, maxX int) []Coord {
	neighbours := make([]Coord, 0)

	if current.Y > 0 { // Top exists
		neighbours = append(neighbours, Coord{current.Y - 1, current.X})
	}

	if current.Y < maxY-1 { // Bottom exists
		neighbours = append(neighbours, Coord{current.Y + 1, current.X})
	}

	if current.X > 0 { // Left exists
		neighbours = append(neighbours, Coord{current.Y, current.X - 1})
	}

	if current.X < maxX-1 { // Right exists
		neighbours = append(neighbours, Coord{current.Y, current.X + 1})
	}

	return neighbours
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
