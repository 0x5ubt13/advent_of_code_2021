package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	lines := getPuzzleInput("./day_8_test.txt")
	var sig, signals, dig, digits []string

	for _, line := range lines {
		sig = append(sig, strings.Split(line, " | ")[0])
		dig = append(dig, strings.Split(line, " | ")[1])
	}



	for _, signal := range sig {
		signals = append(signals, strings.Split(signal, " ")...)
	} 

	for _, digit := range dig {
		digits = append(digits, strings.Split(digit, " ")...)
	} 

	fmt.Println(signals[2], digits[2])
	// for i, line := range signals {
	// 	for j, signal := range line {
	// 		fmt.Printf("Line %d, signal %d -> %s\n", i, j, string(signal))
	// 	}
	// }




}

func getPuzzleInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	chkErr(err)

	lines := strings.Split(string(bytes), "\n")

	// fmt.Println(lines)
	
	return lines
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}