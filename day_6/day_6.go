package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	solve(getPuzzleInput("./day_6_input.txt"))
}

func solve(puzzle []int) {
	school := make(map[int]int)
	shadow := make(map[int]int)

	// Populate map with first values
	for _, v := range puzzle {
		school[v] += 1
	}

	for day := 0; day < 256; day++ {
		shadow[8] += school[0]
		shadow[6] += school[0]
		school[0] = 0

		for i := 8; i > 0; i-- {
			shadow[i-1] += school[i]
		}

		for k, v := range shadow {
			school[k] = v
		}

		// Reset shadow map 
		for i := 0; i < 9; i++ {
			shadow[i] = 0
		}

		// Count the totals
		if day == 79 || day == 255 {
			total := 0

			for _, v := range school {
				total += v
			}

			if day == 79 {
				fmt.Printf("Part 1 -> %d\n", total)
			} else if day == 255 {
				fmt.Printf("Part 2 -> %d\n", total)
			}
		}
	}
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

/* Fun moment with a valuable lesson that I will hardly forget about:

Using the following loop will drain your memory:

for i := 0; i < 79; i++ {
		for j:=0; j < len(school); j++ {
			school[j] -= 1

			if school[j] < 0 {
				school = append(school, 9)
				school[j] = 6
			}	
		}
	}

fmt.Printf("Part 1 -> %d", len(school))


runtime: VirtualAlloc of 54796853248 bytes failed with errno=1455
fatal error: out of memory

runtime stack:
runtime.throw({0x538389, 0xe07f44e000})
        C:/Program Files/Go/src/runtime/panic.go:1198 +0x76
runtime.sysUsed(0xd8e7ecc000, 0xcc225a000)
        C:/Program Files/Go/src/runtime/mem_windows.go:83 +0x1c9
runtime.(*mheap).allocSpan(0x5f3820, 0x66112d, 0x0, 0x1)
        C:/Program Files/Go/src/runtime/mheap.go:1268 +0x3a5
runtime.(*mheap).alloc.func1()
        C:/Program Files/Go/src/runtime/mheap.go:913 +0x69
runtime.systemstack()
        C:/Program Files/Go/src/runtime/asm_amd64.s:383 +0x4e

goroutine 1 [running]:
runtime.systemstack_switch()
        C:/Program Files/Go/src/runtime/asm_amd64.s:350 fp=0xc00051dd30 sp=0xc00051dd28 pc=0x4dada0
runtime.(*mheap).alloc(0x17355fa0f00, 0x17433c50cc8, 0xff, 0x0)
        C:/Program Files/Go/src/runtime/mheap.go:907 +0x73 fp=0xc00051dd80 sp=0xc00051dd30 pc=0x4a3973
runtime.(*mcache).allocLarge(0xc000052000, 0xcc225a000, 0x0, 0x1)
        C:/Program Files/Go/src/runtime/mcache.go:227 +0x89 fp=0xc00051dde0 sp=0xc00051dd80 pc=0x494409
runtime.mallocgc(0xcc225a000, 0x0, 0x0)
        C:/Program Files/Go/src/runtime/malloc.go:1082 +0x5c5 fp=0xc00051de60 sp=0xc00051dde0 pc=0x48bca5
runtime.growslice(0x524b60, {0xceb301e000, 0x1469d5c00, 0x1469d5c00}, 0x1469d5c01)
        C:/Program Files/Go/src/runtime/slice.go:261 +0x4ac fp=0xc00051dec8 sp=0xc00051de60 pc=0x4c54ac
main.main()
        C:/Users/consultant/Documents/GitHub/advent_of_code_2021/day_6/day_6.go:26 +0x1d1 fp=0xc00051df80 sp=0xc00051dec8 pc=0x51c411
runtime.main()
        C:/Program Files/Go/src/runtime/proc.go:255 +0x217 fp=0xc00051dfe0 sp=0xc00051df80 pc=0x4b5ad7
runtime.goexit()
        C:/Program Files/Go/src/runtime/asm_amd64.s:1581 +0x1 fp=0xc00051dfe8 sp=0xc00051dfe0 pc=0x4dd121
exit status 2
*/