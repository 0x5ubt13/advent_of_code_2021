package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	//"strings"
)

func main () {
	data := getLines("./day_3_input.txt")
	partOne(data)
}

func partOne(data []string) {
	commonZeroes := make(map[int]int)
	commonOnes := make(map[int]int)
	mostCommon := make(map[int]int)
	var mc string
	lessCommon := make(map[int]int)
	var lc string

	for i, r := range data {
		// whole binary numbers inside data.
		fmt.Println(i, r) 
		for j, c := range r {
			// binary numbers index inside the numbers
			// string(c) to avoid printing runes
			//fmt.Println(j, string(c)) 
			if string(c) == "0" {
				commonZeroes[j] += 1
			} else {
				commonOnes[j] += 1
			}
		}
	}
	for k := 0; k < len(commonOnes); k++ {
		if commonOnes[k] > commonZeroes[k] {
			mostCommon[k] = 1
			mc += "1"
			lessCommon[k] = 0
			lc += "0"
		} else {
			mostCommon[k] = 0
			mc += "0"
			lessCommon[k] = 1
			lc += "1"
		}
	}
	most, err := strconv.ParseInt(mc, 2, 10)
	chk(err)

	fmt.Println(most, lc)



	
}

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

func chk(e error) {
	if e != nil {
		panic(e)
	}
}