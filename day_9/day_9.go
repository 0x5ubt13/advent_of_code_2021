package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	lines := getPuzzleInput("9.test")

	for _, line := range lines {
		fmt.Println(line)
	}


	// heightmap := makeHeightmap(lines)

	// for k, v := range heightmap {
	// 	fmt.Println(k, v)
	// }
}

func makeHeightmap(lines []int) map[int][]int {
	heightmap := make(map[int][]int, len(lines))
	// for i, row := range lines {}
		

	return heightmap

}

func getPuzzleInput(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")

	intLines := make([]int, 0, 0)
	for _, line := range lines {
		for _, num := range line {
			casting, err := strconv.Atoi(strings.TrimSuffix(string(num), "\r"))
			if err != nil {
				fmt.Println(err)
			}
			intLines = append(intLines, casting)
		}

	}

	return intLines

}