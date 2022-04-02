package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	instructions := getPuzzleInput("./day_5_test.txt")
	
	diagram := drawDiagram(getMaxValue(instructions))
	fmt.Println(diagram)

	// Every even index is "x1, y1" 
	// Every odd  index is "x2, y2"
	for i, l := range instructions {
		if i % 2 == 0 {
			fmt.Printf("x1 = %v, y1 = %v  ", string(l[0]), string(l[2]))
		} else {	
			fmt.Printf("x2 = %v, y2 = %v\n", string(l[0]), string(l[2]))
		}
	}
}

func drawDiagram(max int) map[int]int {
	diagram := make(map[int]int)
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			diagram[i] = 0
		}
	}

	return diagram
}

func getMaxValue(instructions []string) int {
	// Get max value to draw the diagram
	max := 0

	for _, l := range instructions { 
		for _, r := range l {	
			i, _ := strconv.Atoi(string(r))

			if i > max {
				max = i
			}
		}
	}

	return max
}


func getPuzzleInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	chkErr(err)
	
	lines := strings.Split(strings.ReplaceAll(string(bytes), " -> ", " "), "\n")
	instructions := make([]string, 0)

	for _, line := range lines {
		instructions = append(instructions, strings.Split(line, " ")...)
	}

	return instructions
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}