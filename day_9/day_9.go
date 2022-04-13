package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	getPuzzleInput("9.test")
}

func getPuzzleInput(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	clean := strings.Replace(string(bytes), "\r", "")

	lines, err := strconv.Atoi(clean)
	if err != nil {
		panic(err)
	}

	fmt.Println(lines)

}