package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	diagram := drawDiagram()
	diagram.print()
	
	instructions := getPuzzleInput("./day_5_test.txt")

	var x1, x2, y1, y2 int
	
	for i, line := range instructions {
		if i % 2 == 0 {
			// Every even index is "x1, y1" 
			x1, _ = strconv.Atoi(string(line[0]))
			y1, _ = strconv.Atoi(string(line[2]))
			fmt.Printf("x1 = %v, y1 = %v, ", x1, y1)
		} else {
			// Every odd index is "x2, y2"
			x2, _ = strconv.Atoi(string(line[0]))
			y2, _ = strconv.Atoi(string(line[2]))
			fmt.Printf("x2 = %v, y2 = %v\n", x2, y2)
		}

		if i % 2 != 0 && i != 0 {
			// Trigger only every second iteration to make sure we have updated x and y values
			diagram.update(x1, x2, y1, y2)
			diagram.print()
		}
	}

	// ins := [4]int{x1, x2, y1, y2}
}

type Diagram map[int][]int

func (d Diagram) update(x1, x2, y1, y2 int) {
	// X axis
	if x1 == x2 {
		if y1 > y2 {
			// Downwards
			for i := y2; i < y1; i++ {
				d[x1][i] += 1
			}
		} else {
			// Upwards
			for i := y2; i > y1; i-- {
				d[x1][i] += 1
			}
		}

	}
	// Y axis
	if y1 == y2 {
	}
	// Diagonal
}

func (d Diagram) print() {
	for index, row := range d {
		fmt.Println(index, row)
	}
} 

func drawDiagram() Diagram {
	diagram := make(map[int][]int)
	rows := make([]int, 10)
	
	for i := 0; i < 10; i++ {
		diagram[i] = rows
	}

	return diagram
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