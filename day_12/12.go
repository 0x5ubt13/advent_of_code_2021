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
		for _, r := range line {
			fmt.Println(r)
		}

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

	// 2: Create Caves
	Caves := make([]Cave, 0)
	for _, c := range uniqueCaves {
		Caves = append(Caves, createCave(c, connections))
	}

	// 3: Find all paths (breadth-first search)
	paths := make(map[int][]string, 0)
	queue := []Path{Path{[]string{"start"}, map[string]bool{}}}

	// while queue not empty
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:] 

		for i, cave := range Caves {
			fmt.Println(i, cave.Name)
			
			for _, ca := range Caves {
				if ca.Name == c {
					fmt.Println(ca)
				}
			}

			j, match := strInStrSlice(cave.Links, c)
			if match == true {
				// connection match

				fmt.Printf("match: index=%d | cave.Links=%s\n", j, c)




			}
			
		}
	}

	fmt.Println(paths)


}

type Path struct {
	Path 	[]string
	Visited map[string]bool
}

type Cave struct {
	Name 	string
	IsBig 	bool
	Links 	[]string
}

func copyMap(a map[string]bool) map[string]bool {
	out := make(map[string]bool)
	for k, v := range a {
		out[k] = v
	}

	return out
}

func createCave(newCave string, connections []string) Cave {
	var cave Cave
	links := make([]string, 0)

	// Name
	cave.Name = newCave

	// IsBig
	if strings.ToUpper(newCave) == newCave {
		cave.IsBig = true
	} else {
		cave.IsBig = false
	}

	// Links
	for _, l := range connections {
		link := strings.Split(strings.TrimSuffix(l, "\r"), "-")

		if cave.Name == link[0] {
			links = append(links, link[1])
		} else if cave.Name == link[1] {
			links = append(links, link[0])
		} else {
			continue
		}
	}

	uniqueLinks := uniqueNonEmptyElementsOf(links)
	for _, uniqueLink := range uniqueLinks {
		cave.Links = append(cave.Links, uniqueLink)
	}

	// fmt.Println(cave)

	return cave
}

func strInStrSlice(sl []string, s string) (int, bool) {
	var x int
	for i, ss := range sl {
		if s == ss {
			x = i
			return x, true
		}
	}
	
	return x, false
}

func uniqueNonEmptyElementsOf(s []string) []string {
	uniqueCheck := make(map[string]bool, len(s))
	uniqueSlice := make([]string, len(uniqueCheck))
		for _, elem := range s {
			elem = strings.TrimSuffix(elem, "\r")
			if len(elem) != 0 {
				if !uniqueCheck[elem] {
					uniqueSlice = append(uniqueSlice, elem)
					uniqueCheck[elem] = true
				}
			}
		}
	
		return uniqueSlice
}


func getPuzzleInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(string(bytes), "\n")
}