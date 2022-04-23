package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	connections := getPuzzleInput("./12.test")
	// paths := make(map[int][]string, 0)
	caves := []string{"start"}

	/*
	Your goal -> find the number of distinct paths that start at start, end at end, and don't visit small caves + than 1. 
	There are two types of caves: 
			big caves (written in uppercase, like A) can be visited multiple times 
			small caves (written in lowercase, like b), only visited once 
	*/

	// 1: get connections and create paths
	for i, line := range connections {
		fmt.Println(i, line)

		newCaves := strings.Split(line, "-")

		for _, newCave := range newCaves {			
			caves = append(caves, newCave)
		}
	}


	fmt.Println("All caves:")
	uniqueCaves := uniqueNonEmptyElementsOf(caves)
	for _, c := range uniqueCaves {
		fmt.Println(c)
	}

	// fmt.Println("Starting points:")

	// 2: explore connections
	//		first, note down every start-something1 or start-something2
	//		then,
	//			 add something1-something3 to the one ending in something1
	//			 add something2-something3 to the one ending in something2
	//		finally, add something3-end

}

func uniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
		us := make([]string, len(unique))
		for _, elem := range s {
			if len(elem) != 0 {
				if !unique[elem] {
					us = append(us, elem)
					unique[elem] = true
				}
			}
		}
	
		return us
  
  }

func removeStrFromSlice(s []string, r string) []string {
	for k, v := range s {
		if r == v {
			if s[len(s)-1] == r {
				s = s[:k]
			} else {
				s = append(s[:k], s[k+1:]...)
			}
		}
	}

	return s
}

func stringInSlice(y []string, x string) bool {
	for _, s := range y {
		if s == x {
			return true
		}
	}

	return false
}

func getPuzzleInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(string(bytes), "\n")
}