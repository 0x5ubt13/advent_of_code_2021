package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	caves := getPuzzleInput("./12.test")
	paths := make(map[int][]string, 0)

	for i, line := range caves {
		newPath := strings.Split(line, "-")
		start := newPath[0]
		end := newPath[1]
		paths[i] = append(paths[i], start, end)

		fmt.Println(start, end)
	}


}

func getPuzzleInput(filename string) []string {
	bytes, _ := ioutil.ReadFile(filename)

	return strings.Split(string(bytes), "\n")
}