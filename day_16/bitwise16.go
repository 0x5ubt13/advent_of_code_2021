package main

import "fmt"

func main() {

	bits := []byte{
		1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0,
	}

	fmt.Println(readBits(bits, 0, 3))
}

func readNumber(data []byte, startpos int) int64 {
	var out int64

	part := readBits(data, startpos, 5)
	out <<= 4
	out |= int64(part & 0x0f)
}

func readBits(data []byte, startpos, count int) int64 {
	var out int64
	for i, b := range data[startpos:startpos+count] {
		fmt.Printf("Reading bits, iter %d: b = %d\n", i, b)
		fmt.Printf("Current out:\nDecimal: %d\t Binary: %b\n", out, out)
		out <<= 1
		fmt.Printf("New out after bitshifting left 1:\nDecimal: %d\t Binary: %b\n", out, out)
		out |= int64(b)
		fmt.Printf("New out after out or-equals int64(b):\nDecimal: %d\t Binary: %b\n\n\n", out, out)
	}

	return out
}