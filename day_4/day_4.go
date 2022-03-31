package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	rawBytes := getInput("day_4_input.txt")
	lines := strings.Split(string(rawBytes), "\n")
	partOne, partTwo := solve(lines)
	fmt.Printf("Part 1 -> %d\n", partOne)
	fmt.Printf("Part 2 -> %d\n", partTwo)
}

func solve(lines []string) (int, int) {
	var result1, result2 int
	numbersPool := strings.Split(lines[0], ",")
	drawnNumbers := make([]int, 0)
	winningBoardsTracker := make([]int, 0)
	newRound(drawnNumbers, numbersPool)

	// 0. create boards -> map int:[]int -> 1: board, 2: board, 3: board etc
	playingBoards := playingBoards(lines)

	// Starting the game, main loop
	for {
		// 1. draw new number
		drawnNumbers = newRound(drawnNumbers, numbersPool)
		if drawnNumbers[len(drawnNumbers)-1] == 999 {
			break
		}
		fmt.Println("Numbers drawn:", drawnNumbers)
		
		// 2. loop drawn nums for every board possible. Check against shadow boolean board
		for board, numbers := range playingBoards {
			// Part 2 -> check that the board has already won, skip it if so
			chk := numInSlice(board, winningBoardsTracker)
			if chk == true {
				continue
			} 

			// Create new shadow board for every new board
			fmt.Printf("Board number #%d. Creating new shadow board.\n", board)
			shadowBoard := newShadowBoard()

			for row, column := range numbers {
				for col, number := range column {
					// fmt.Printf("row %v, column %v, number %v\n", row, col, number)
					for _, drawnNum := range drawnNumbers {
						if drawnNum == number {
							shadowBoard[row][col] = true
						}
					}
				}
			}

			for _, row := range shadowBoard {
				fmt.Println(row)
			} 
			
			// Getting some winning boards. Save the first for part 1 and continue until hitting the last one for part 2
			part1Endgame := checkBoard(shadowBoard)

			if part1Endgame == true {
				winningBoardsTracker = append(winningBoardsTracker, board)
				lastCall := drawnNumbers[len(drawnNumbers)-1]

				if result1 == 0 {	
					result1 = calculateResult(shadowBoard, playingBoards[board], lastCall)
				}
				
				if len(winningBoardsTracker) == len(playingBoards) {
					result2 = calculateResult(shadowBoard, playingBoards[board], lastCall)
				}

				part1Endgame = false
			}

			if result2 != 0 {
				break
			}
		}

		if result2 != 0 {
			break
		}
		
	} 

	return result1, result2
}

func calculateResult(shadowBoard [5][5]bool, playingBoard [5][5]int, lastCall int ) int {
	var result int
	for row, cols := range shadowBoard {
		for col, value := range cols {
			if value == false {
				result += playingBoard[row][col]
			}
		}
	}

	return result * lastCall
}

func newRound(drawnNumbers []int, numbersPool []string) []int {
	newNumber := drawNewNumber(drawnNumbers, numbersPool)
	drawnNumbers = append(drawnNumbers, newNumber)
	return drawnNumbers
}

func drawNewNumber(drawnNumbers []int, numbersLine []string) int {
	for _, n := range numbersLine {
		n, err := strconv.Atoi(strings.TrimSuffix(n, "\r"))
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
	}
	
	return false
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

func checkBoard(board [5][5]bool) bool {
	winner := false
	var row, col int

	for i := 0; i < 5; i++ {
		//reset loop
		row = 0
		col = 0

		for j := 0; j < 5; j++ {
			// horizontal check
			if board[i][j] == true{ 
				row++
			}

			// vertical check
			if board[j][i] == true {
				col++
			}
		}

		if row == 5 || col == 5 {
			winner = true
		}
	}

	return winner
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