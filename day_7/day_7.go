package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	solve(getPuzzleInput("./day_7_input.txt"))
}

func solve(crabs []int) {
	var fuel, premiumFuel, min, premiumMin, steps int
	petrolStation := make(map[int]int, len(crabs))
	premiumPetrolStation := make(map[int]int, len(crabs))

	for i := 0; i < len(crabs) + 1; i++ {
		fuel = 0
		premiumFuel = 0

		for _, crab := range crabs {
			steps = 0

			if i > crab {
				steps = i - crab
				fuel += steps
				premiumFuel += steps
			} else if {
				steps := crab - i
				fuel += steps
				premiumFuel += steps
			}
	
			for s := 0; s < steps + 1; s++{
				premiumFuel += s
			}
				
		}
		petrolStation[i] = fuel
		premiumPetrolStation[i] = premiumFuel
	}
	fmt.Println(premiumPetrolStation)

	min = petrolStation[0]
	for _, v := range petrolStation {
		if v < min {
			min = v
		}
	}

	premiumMin = premiumPetrolStation[0]
	for _, w := range premiumPetrolStation {
		if w < premiumMin {
			premiumMin = w
		}
	}

	fmt.Printf("Part 1 -> %d\n", min)
	fmt.Printf("Part 2 -> %d\n", premiumMin)

}

func getPuzzleInput(filename string) []int {
	var crabs []int
	
	bytes, err := ioutil.ReadFile(filename)
	chkErr(err)
	
	lines := strings.Split(string(bytes), ",")
	
	for _, l := range lines {
		a, err := strconv.Atoi(l)
		chkErr(err)
		crabs = append(crabs, a)
	}

	return crabs

}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}