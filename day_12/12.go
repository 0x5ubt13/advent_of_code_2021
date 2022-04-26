package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "Puzzle day 12")
	partOne()
	partTwo()
}

func partOne() {
	connections := getPuzzleInput("./12.in")
	caves := []string{"start"}

	/*
	Your goal -> find the number of distinct paths that start at start, end at end, and don't visit small caves + than 1. 
	There are two types of caves: 
			big caves (written in uppercase, like A) can be visited multiple times 
			small caves (written in lowercase, like b), only visited once 
	*/

	// 1: get connections and create paths
	for _, line := range connections {
		// for _, r := range line {
		// 	fmt.Println(r)
		// }

		newCaves := strings.Split(line, "-")

		for _, newCave := range newCaves {			
			caves = append(caves, newCave)
		}
	}

	// fmt.Println("All caves:")
	uniqueCaves := uniqueNonEmptyElementsOf(caves)
	// for _, c := range uniqueCaves {
	// 	fmt.Println(c)
	// }

	// 2: Create Caves
	Caves := make(map[string]Cave, 0)
	for _, c := range uniqueCaves {
		Caves[c] = createCave(c, connections)
	}

	// fmt.Println(Caves["start"])
	
	//3: Find all paths (breadth-first search)
	queue := []Path{{[]string{"start"}, make(map[string]bool, 0)}}
	paths := make([]string, 0)

	//while queue not empty
	for len(queue) > 0 {
		// Pop the first element of the queue (first iteration -> start)
		cur := queue[0]
		queue = queue[1:] 

		// assign cave to the struct Cave being visited (first iteration -> "{name:start isbig:false links:[A b]}")
		cave := Caves[cur.Path[len(cur.Path)-1]]

		// If reached the end, stop
		if cave.Name == "end" {
			paths = append(paths, strings.Join(cur.Path, ","))
		}

		// Copying the map is required to avoid overwriting the underlying shared map
		newVisited := copyMap(cur.Visited)
		if !cave.IsBig {
			newVisited[cave.Name] = true
		}

		// Visit all the links the current cave is linked to
		for _, kave := range cave.Links {
			// Original name for a cave, I know :) 
			// running out of valid, relevant names here

			// If in the current iteration of the queue the cave "kave" was visited, abort
			if cur.Visited[kave] {
				continue
			}

			// copy the path to avoid overwriting the underlying shared array
			newPath := make([]string, len(cur.Path))
			copy(newPath, cur.Path)

			newPath = append(newPath, kave)

			queue = append(queue, Path{newPath, newVisited})
		}
	}

	// sort.Strings(paths)
	// for _, path := range paths {
	// 	fmt.Println(path)
	// }

	fmt.Printf("Part 1 -> %d different paths\n", len(paths))

}

func partTwo() {
	connections := getPuzzleInput("./12.in")
	caves := []string{"start"}

	for _, line := range connections {
		newCaves := strings.Split(line, "-")

		for _, newCave := range newCaves {			
			caves = append(caves, newCave)
		}
	}

	uniqueCaves := uniqueNonEmptyElementsOf(caves)

	Caves := make(map[string]Cave, 0)
	for _, c := range uniqueCaves {
		Caves[c] = createCave(c, connections)
	}

	queue := []PathTwo{{[]string{"start"}, map[string]bool{}, false}}
	paths := make([]string, 0)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:] 

		cave := Caves[cur.Path[len(cur.Path)-1]]

		if cave.Name == "end" {
			paths = append(paths, strings.Join(cur.Path, ","))
			continue
		}

		newVisited := copyMap(cur.Visited)
		if !cave.IsBig {
			newVisited[cave.Name] = true
		}

		for _, kave := range cave.Links {
			newSmall := cur.SmallVisitedTwice

			if cur.Visited[kave] { 
				if kave == "start" || newSmall {
					continue
				} else {
					newSmall = true
				}
			}

			// copy the path to avoid overwriting the underlying shared array
			newPath := make([]string, len(cur.Path))
			copy(newPath, cur.Path)

			newPath = append(newPath, kave)

			queue = append(queue, PathTwo{newPath, newVisited, newSmall})
		}
	}

	// for _, path := range paths {
	// 	fmt.Println(path)
	// }

	fmt.Printf("Part 2 -> %d different paths\n", len(paths))

}

type Path struct {
	Path 	[]string
	Visited map[string]bool
}

type PathTwo struct {
	Path 	[]string
	Visited map[string]bool
	SmallVisitedTwice bool
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

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %v seconds", name, elapsed.Seconds())
}