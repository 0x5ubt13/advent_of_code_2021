package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	//"strings"
)

func main () {
	data := getLines("./day_3_input.txt")

	fmt.Print(data)

	for i, r := range data {
		fmt.Println(i, r)
		for j, c := range r {
			fmt.Println(j, string(c))
		}
	}
}

func getLines(filename string) []string {
	f, err := os.Open(filename)
	chk(err)
	defer f.Close()

	data := make([]string, 0)

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		data = append(data, string(line))
	}

	return data
}

func chk(e error) {
	if e != nil {
		panic(e)
	}
}