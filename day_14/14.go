package main

import(
	//"fmt"
	"log"
	"os"
	"bufio"
	"strconv"

)

func main() {
	getInput("./14_test.txt")
}

func getInput(filename string) []string {
	lines := make()

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line, err2 := scanner.Text()
		if err2 != nil {
			log.Fatal(err2)
		}

		lines = append(lines, line)
	}
	return lines
}