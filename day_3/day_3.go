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
	o2 := make([]string, len(data))
	co2 := make([]string, len(data))

	// Duplicate data for the condition on the first iteration to be met
	copy(o2, data)
	copy(co2, data)
	fmt.Println(o2, co2)
	

	// Find the most common bit and eliminate the rest
	for i := 0; i < 12; i++ {
		// Reset the following vars every iteration
		var zeroes, ones int

		// Check which is the most common, either 0 or 1
		for _, n := range data {
			for j, r := range n {
				if j == i {
					if string(r) == "0" {
						zeroes += 1		
					} else {
						ones += 1
					}
				}
			}
		}

		//fmt.Println(zeroes, ones)

		// Save only those who fit the criteria
		for _, n := range data {
			if ones >= zeroes {
				//fmt.Println(string(n[i]))
				check := stringInSlice(n, o2)
				fmt.Println(check)
				if string(n[i]) == "1" && check == true {
					
				} else{
					removeFromSlice(o2, n)
				}
			} else {
				// Save zeroes
				check := stringInSlice(n, co2)
				if string(n[i]) == "0" && check == true {
					
				} else {
					removeFromSlice(co2, n)
				}
			}
		}
	}

		fmt.Println(len(o2), len(co2))

	fmt.Println(o2, co2)

	fmt.Println("Part 2 solution ->")
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

// Utility function similar to Pythons "if x in list"
func stringInSlice(x string, slice []string) bool {
	for _, b := range slice {
		if b == x {
			return true
		}
	}
	return false
}



func removeFromSlice(s []string, r string) []string {
    for i, v := range s {
        if v == r {
            return append(s[:i], s[i+1:]...)
        }
    }
    return s
}

// Condensing error handling
func chk(e error) {
	if e != nil {
		panic(e)
	}
}
