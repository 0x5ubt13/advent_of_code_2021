package main

import (
	"fmt"
)

func main() {
	target := createTargetArea(88, 125, -103, -157)
	target.test()

	finalY := aim(target)
	fmt.Println("Y =", finalY)

}

func aim(target *Target) int {
	velocityY := -20
	velocityX := 50
	i := 0

	for velocityX > 0 {
		try, steps := shootProbe(velocityY, velocityX, target)
		if try == true {
			fmt.Println("HIT!", try, steps, velocityY)
			return velocityY
		}

		velocityY--
		if i % 2 == 0 {
			velocityX++
		}
	} 

	return 0
}

func shootProbe(y, x int, target *Target) (bool, int) {
	// Aceleration loses x-1 and x-1 per step. X can't be -1.
	step := 0
	for {
		if x < 0 {
			return false, step
		}

		if x > target.X["Max"]  {
			return false, step
		}

		if y < target.Y["Max"] {
			return false, step
		}

		fmt.Printf("Step %d, x = %d, y = %d ", step, x, y)
		// Target hit
		if x > target.X["Min"] && x < target.X["Max"] {
			if y > target.Y["Min"] && y < target.Y["Max"] {
				fmt.Println("HIT!")
				return true, step
			}
		} 

		// End of step
		step++
		
		y--
		fmt.Printf("Miss.\n")
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

	for y := minY; y > maxY ; y-- {
		target.Area["Y"] = append(target.Area["Y"], y)
	}

	for x := minX; x < maxX ; x++ {
		target.Area["X"] = append(target.Area["X"], x)
	}

	return target

}

type Target struct {
	X 		map[string]int
	Y 		map[string]int
	Area 	map[string][]int
}