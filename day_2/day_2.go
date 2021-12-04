package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	m := loadInput("./day_2/day_2_input.txt")

	directions := map[string]int{
		"horizontal" : 	0,
		"depth" : 		0,
		"depthTwo":		0,
		"aim":			0,
	}

	for _, d := range m {
		for direction, instruction := range d {
			switch direction {
			case "forward":
				directions["horizontal"] += instruction
				directions["depthTwo"] += directions["aim"] * instruction
			case "down":
				directions["depth"] += instruction
				directions["aim"] += instruction
			case "up":
				directions["depth"] -= instruction
				directions["aim"] -= instruction
			}
		}
	}

	fmt.Printf("Part 1 %d | Part 2 %d\n", directions["horizontal"] * directions["depth"], directions["horizontal"] * directions["depthTwo"])

}

func loadInput(filename string) []map[string]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var directions []map[string]int
	for _, val := range lines {
		tempSlice := strings.Fields(val)
		key := tempSlice[0]
		value, err2 := strconv.Atoi(tempSlice[1])
		if err2 != nil {
			log.Fatalf("Error: %v", err2)
		}
		m := map[string]int{key: value}
		directions = append(directions, m)
	}

	return directions
}


