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
	solutionOne := partOne(lines)
	fmt.Println(solutionOne)
}

func partOne(lines []string) string {
	numbersPool := strings.Split(lines[0], ",")
	drawnNumbers := make([]int, 0)
	newRound(drawnNumbers, numbersPool)

	// 0. create boards -> map int:[]int -> 1: board, 2: board, 3: board etc
	playingBoards := playingBoards(lines)
	// for i, b := range playingBoards {
	// 	fmt.Println("Board number",i, ":")
	// 	for _, board := range b {
	// 		fmt.Println("\t\t", board)
	// 	}
	// }
	
	// Starting the game, main loop
	for {
		// 1. draw new number
		drawnNumbers = newRound(drawnNumbers, numbersPool)
		if drawnNumbers[len(drawnNumbers)-1] == 999 {
			return "Exhausted numbers pool. Try harder."
		}
		fmt.Println("Numbers drawn:", drawnNumbers)
		
		// 2. loop drawn nums for every board possible. Check against shadow boolean board
		for i, board := range playingBoards {
			fmt.Printf("Board number #%d\n", i)
			for row, column := range board {
				for col, number := range column {
					//test fmt.Printf("row %v, column %v, number %v\n", row, col, number)
					
				}
			}
		}
	}
}

func partTwo() {

}

func newRound(drawnNumbers []int, numbersPool []string) []int {
	newNumber := drawNewNumber(drawnNumbers, numbersPool)
	drawnNumbers = append(drawnNumbers, newNumber)
	return drawnNumbers
}

func drawNewNumber(drawnNumbers []int, numbersLine []string) int {
	for _, n := range numbersLine {
		n, err := strconv.Atoi(n)
		if err != nil{fmt.Println(err)}
		if numInSlice(n, drawnNumbers) == false {
			return n
		}
	}
	return 999
} 

func numInSlice(x int, y []int) bool {
	for _, n := range y {
		if x == n {
			return true
		}
	}; return false
}

func newBoard() [5][5]int {
	var newBoard [5][5]int = [5][5]int{
		{0, 0, 0, 0, 0}, 
		{0, 0, 0, 0, 0}, 
		{0, 0, 0, 0, 0}, 
		{0, 0, 0, 0, 0}, 
		{0, 0, 0, 0, 0}}

	return newBoard
}

func newShadowBoard() [5][5]bool {
	var newShadowBoard [5][5]bool = [5][5]bool{
		{false, false, false, false, false}, 
		{false, false, false, false, false}, 
		{false, false, false, false, false}, 
		{false, false, false, false, false}, 
		{false, false, false, false, false}}

	return newShadowBoard
}

func checkBoard(drawnNumbers, board []int) bool {
	var winner bool
	
	if winner == true {
		fmt.Println("WE HAVE A WINNER!!!")
		return true
	} else {
		return false
	}
}

func playingBoards(lines []string) map[int][5][5]int {
	board := newBoard()
	rowPointer := 0
	boardPointer := 0
	playingBoards := make(map[int][5][5]int)

	for _, line := range lines[2:] {	
		// Populate new board
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

			// New row please
			rowPointer++

			// Once fully populated, save board and start over
			if rowPointer == 5 {
				playingBoards[boardPointer] = board
				boardPointer++
			}
		}	
	}
	return playingBoards
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
