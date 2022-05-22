package main

import (
	"fmt"
)

func main() {
	// target := createTargetArea(20, 30, -10, -5) // Test input
	target := createTargetArea(88, 125, -157, -103) // Puzzle input
	// target.test()

	targetHitsCoordinates := autoAim(target)
	// fmt.Printf("%+v\n", targetHitsCoordinates)

	part1 := 0
	for _, coord := range targetHitsCoordinates {
		if coord.maxY > part1 {
			part1 = coord.maxY
		}
	}
	fmt.Printf("Part 1 -> %d\n", part1)
 	fmt.Printf("Part 2 -> %d\n", len(targetHitsCoordinates))
}

func autoAim(target *Target) ([]Coord) {
	targetHitsCoordinates := make([]Coord, 0)

	for x := 0; x < -target.Y["Min"]; x++ {
		for y := target.Y["Min"]; y < -target.Y["Min"]; y++ {
			hit, maxY := shootProbe(x, y, target)
			if hit == true {
				targetHitsCoordinates = append(targetHitsCoordinates, Coord{x, y, maxY})
			}
		}
	}
	
	return targetHitsCoordinates
}

func shootProbe(x, y int, target *Target) (bool, int) {
	// Aceleration loses x-1 and x-1 per step. X can't be -1.
	maxY := 0
	probeCoords := Coord{0,0,0}
	for {
		probeCoords.X += x
		probeCoords.Y += y
		if probeCoords.Y > maxY {
			maxY = probeCoords.Y
		}
		// fmt.Printf("Step %d, speedX = %d, speedY = %d\nProbeX = %d, ProbeY = %d\n", step, x, y, probeCoords.X, probeCoords.Y)

		// Miss returns false, hit returns true
		if x > target.X["Max"]  {
			return false, 0
		}

		if y < target.Y["Min"] {
			return false, 0
		}

		// Target hit
		if probeCoords.X >= target.X["Min"] && probeCoords.X <= target.X["Max"] && probeCoords.Y >= target.Y["Min"] && probeCoords.Y <= target.Y["Max"] {
			return true, maxY
		} 

		// End of step
		if x > 0 {
			x--
		}
		y--
	}
} 

func (target *Target) test() {
	// Testing whether target has been implemented correctly
	fmt.Println(*target)

	fmt.Println(target.X)
	fmt.Println(target.Y)
	
	fmt.Println(target.X["Min"])
	fmt.Println(target.Y["Min"])

	fmt.Printf("%+v\n", *target)
}

func createTargetArea(minX, maxX, minY, maxY int) *Target {
	target := new(Target)
	
	target.X = 		make(map[string]int)
	target.Y = 		make(map[string]int)
	target.Area =	make(map[string][]int)  

	target.X["Min"] = minX
	target.X["Max"] = maxX
	target.Y["Min"] = minY
	target.Y["Max"] = maxY
	target.Area["Y"] = make([]int, 0)
	target.Area["X"] = make([]int, 0)

	for y := minY; y < maxY ; y++ {
		target.Area["Y"] = append(target.Area["Y"], y)
	}

	for x := minX; x < maxX ; x++ {
		target.Area["X"] = append(target.Area["X"], x)
	}

	return target
}

type Coord struct {
	X, Y, maxY int
}

type Target struct {
	X 		map[string]int
	Y 		map[string]int
	Area 	map[string][]int
}