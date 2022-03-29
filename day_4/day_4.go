package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	rawBytes := getInput("day_4_test_input.txt")
	lines := strings.Split(string(rawBytes), "\n")
	partOne(lines)
}

func partOne(lines []string) {
	board := newBoard()
	rowPointer := 0

	// Draw numbers 
	// drawnNumbers := strings.Split(lines[0], ",")
	// for _, num := range drawnNumbers {
	// }
	//fmt.Println("-------------------")


	// Scan boards line by line and compare to drawn numbers
	for _, line := range lines[2:] {

		if strings.TrimSpace(line) == "" {
			// White line, create new board
			board = newBoard()
			rowPointer = 0
		} else {
			// Create nums row
			row := []int{}
			numbers := strings.Split(line, " ")
			for _, num := range numbers {
				if num != "" {
					newNum, err := strconv.Atoi(strings.TrimSuffix(num, "\r"))
					if err != nil {
						fmt.Println(err)
					}
					row = append(row, newNum)
				}
			}

			// Fill up the board with the new row
			for index, number := range row {
				board[rowPointer][index] = number
			}
			
			// Finally, compare drawn numbers to the boards and check if there is a matching line
			for c, r := range board { 
				fmt.Println(c, r)
			}

			rowPointer++
		}

	}
	
}

func partTwo() {

}

func newBoard() [5][5]int {
	var newBoard [5][5]int = [5][5]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}

	for i := 0; i < len(newBoard); i++ {
		fmt.Println(newBoard[i])
	}

	return newBoard
}

func getInput(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	err_chk(err)
	return bytes
}

func err_chk(err error) {
	if err != nil {
		panic(err)
	}
}
