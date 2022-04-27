package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main () {
	coords, folds := getPuzzleInput("./13.test")

	for i, c := range coords {
		fmt.Println(i, c)
	}

	for i, f := range folds {
		fmt.Println(i, f)
	}

}

type Coord struct {
	X int
	Y int
}

type Fold struct {
	Axis string
	Position int
}

func getPuzzleInput(fn string) ([]Coord, []Fold) {
	bytes, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	coords := make([]Coord, 0)
	folds := make([]Fold, 0)

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		var newCoord Coord
		var newFold Fold
		
		coordinates := strings.Split(line, ",")
		newCoord.X, err = strconv.Atoi(coordinates[0])
		if err != nil {
			if len(line) > 1 {
				for _, ch := range line {
					if ch == 'y' {
						newFold.Axis = "y"
					} else if ch == 'x' {
						newFold.Axis = "x"
					} else if int(ch) > 1 {
						newFold.Position = int(ch) - 48
					}
				}

				folds = append(folds, newFold)
			}
			continue
		} 
		newCoord.Y, _ = strconv.Atoi(coordinates[1])
		
		coords = append(coords, newCoord)
	}

	return coords, folds
}

