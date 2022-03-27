package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	//"strings"
)

func main() {
	data := getLines("./day_3_input.txt")
	// partOne(data)
	partTwo(data)
}

// Part 1 solution
func partOne(data []string) {
	commonZeroes := make(map[int]int)
	commonOnes := make(map[int]int)
	var mostCommon, leastCommon string

	// Parse most common bits bit by bit
	for _, r := range data {
		// whole binary numbers inside data.
		//fmt.Println(i, r) //test
		for j, c := range r {
			// binary numbers index inside the numbers
			// string(c) to avoid printing runes
			//fmt.Println(j, string(c)) //test
			if string(c) == "0" {
				commonZeroes[j] += 1
			} else {
				commonOnes[j] += 1
			}
		}
	}

	for k := 0; k < len(commonOnes); k++ {
		if commonOnes[k] > commonZeroes[k] {
			mostCommon += "1"
			leastCommon += "0"
		} else {
			mostCommon += "0"
			leastCommon += "1"
		}
	}

	most, err := strconv.ParseInt(mostCommon, 2, 64)
	chk(err)

	least, err := strconv.ParseInt(leastCommon, 2, 64)
	chk(err)

	fmt.Println("Part 1 solution ->", most*least)
}

// Part 2 solution
func partTwo(data []string) {
	oxygen := make([]string, 0) 
	co2 := make([]string, 0)
	commonOnes := make(map[int]int)
	commonZeroes := make(map[int]int)

	// Making working copies of original data
	for _, v := range data {
		oxygen = append(oxygen, v)
		co2 = append(co2, v)
	}

	// Oxygen values (common 1)
	// Find the most common bit and eliminate the rest
	for i := 0; i < 12; i++ {
		for _, n := range oxygen {
			for j, r := range n {
				if string(r) == "0" {
					commonZeroes[j] += 1
				} else {
					commonOnes[j] += 1
				}
			}
		}
		fmt.Println(commonOnes)
	}
}

// Getting input
func getLines(filename string) []string {
	f, err := os.Open(filename)
	chk(err)
	defer f.Close()

	data := make([]string, 0)

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		data = append(data, string(line))
	}

	return data
}

// Condensing error handling
func chk(e error) {
	if e != nil {
		panic(e)
	}
}
