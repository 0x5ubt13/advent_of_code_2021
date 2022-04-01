package main 

import(
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	getPuzzleInput("./day_5_test.txt")
	// instructions := make(map[int][]int)

	// for _, l := range lines {
	// 	fmt.Println(l)
	// }
}

func getPuzzleInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	chkErr(err)
	
	lines := strings.Split(strings.ReplaceAll(string(bytes), " -> ", " "), "\n")
	ins := make([]string, 0)

	for _, line := range lines {
		ins = append(ins, strings.Split(line, " ")...)
	}

	for i, l := range ins {
		if i % 2 == 0 {
			fmt.Printf("x1 = %v, y1 = %v  ", string(l[0]), string(l[2]))
		} else {	
			fmt.Printf("  x2 = %v, y2 = %v\n", string(l[0]), string(l[2]))
		}
	}

	return lines
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}