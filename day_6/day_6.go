package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)
// In this example, after 18 days, there are a total of 26 fish. After 80 days, there would be a total of 5934.
func main() {
	lanternfishes := getPuzzleInput("./day_6_test.txt")
	
	school := make(map[int]int)

	for _, l := range lanternfishes {
		school[l] += 1
	}

	fmt.Println(school)

	for i := 0; i > 18; i++ {
		shadowmap := make(map[int]int)

	
		shadowmap[8] = school[0]
		shadowmap[6] = school[0]

		j := 6
		for i:=7; i>1; i--{
			shadowmap[i] = school[j]
			j--
		}

		for k, v := range shadowmap {
			school[k] = v
		}
	}

	fmt.Println(school)

	total := 0
	for k, _ := range school {
		total += school[k]
	}
	fmt.Println(total)

}

func tick(school map[int]int) map[int]int {
	shadowmap := make(map[int]int)
	
	shadowmap[8] = school[0]
	shadowmap[6] = school[0]

	j := 6
	for i:=7; i>1; i--{
		shadowmap[i] = j
		j--
	}

	fmt.Println(shadowmap)

	return shadowmap



}

func getPuzzleInput(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	chkErr(err)

	var lanternfishes []int
	for _, l := range bytes {
		if string(l) != "," {
			lan, err := strconv.Atoi(string(l))
			chkErr(err)
			lanternfishes = append(lanternfishes, lan)
		}
	}

	return lanternfishes
}

func chkErr(err error){
	if err != nil {
		panic(err)
	}
}