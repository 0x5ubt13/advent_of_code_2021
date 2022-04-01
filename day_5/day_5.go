package main 

import(
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	lines := getPuzzleInput("./day_5_test.txt")
	for _, l := range lines {
		fmt.Println(l)
	}
}

func getPuzzleInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	chkErr(err)
	lines := strings.Split(string(bytes), "\n")
	return lines
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}