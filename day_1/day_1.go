package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Part 1
	increases := 0
	depths := getInput("./day_1/day_1_input.txt")
	for i := 1; i <= len(depths)-1; i++ {
		if depths[i] > depths[i-1] {
			increases += 1
		}
	}
	fmt.Printf("Part 1: %v", increases)

	// Part 2
	partTwoIncreases := 0
	partTwoDepths := make([]int64, 0)
	var quickResult int64
	for i := 0; i <= len(depths)-3; i++ {
		quickResult = depths[i] + depths[i+1] + depths[i+2]
		partTwoDepths = append(partTwoDepths, quickResult)
	}

	for j := 1; j <= len(partTwoDepths)-1; j++ {
		if partTwoDepths[j] > partTwoDepths[j-1] {
			partTwoIncreases += 1
		}
	}
	fmt.Printf("Part 2: %v", partTwoIncreases)

}

func getInput(filename string) []int64 {
	lines := make([]int64, 0)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		intLine, err2 := strconv.ParseInt(scanner.Text(), 10, 64)
		if err2 != nil {
			log.Fatal(err2)
		}

		lines = append(lines, intLine)
	}
	return lines
}
