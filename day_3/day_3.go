package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	data := getLines("./day_3_input.txt")
	partOne(data)
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
		for j, c := range r {
			// binary numbers index inside the numbers - string(c) to avoid printing runes
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
	var oxygen, dioxide int64
	var common0, common1 []string
	o2 := data
	co2 := data

	// Find the most common bit and eliminate the rest
	// Oxygen
	column := 0
	for {
		// Separate, then check which is the most common, either 0 or 1
		for _, n := range o2 {
			if string(n[column]) == "1" {
				common1 = append(common1, n)
			} else {
				common0 = append(common0, n)
			}
		}

		column++

		// Save only those who fit the criteria
		if len(common1) >= len(common0) {
			// Save ones
			o2 = common1
		} else {
			// Save zeroes
			o2 = common0
		}
		
		// Reset
		common0 = common0[:0]
		common1 = common1[:0]

		if len(o2) == 1 {
			oxygen, _ = strconv.ParseInt(o2[0], 2, 64)
			break
		}
	}
	
	// Dioxide -> repeat
	column = 0 
	for {
		for _, n := range co2 {
			if string(n[column]) == "1" {
				common1 = append(common1, n)
			} else {
				common0 = append(common0, n)
			}
		}
		
		column++
			
		if len(common1) >= len(common0) {
			co2 = common0
		} else {
			co2 = common1
		}
			
		if len(co2) == 1 {
			dioxide, _ = strconv.ParseInt(co2[0], 2, 64)
			break
		}
			
		// Reset
		common0 = common0[:0]
		common1 = common1[:0]
	}
			
	fmt.Println("Part 2 solution ->", oxygen * dioxide)
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
