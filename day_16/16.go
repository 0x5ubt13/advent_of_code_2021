package main

import (
	"fmt"
	"io/ioutil"
	// "strings"
	"strconv"
)

func main() {
	// data := getPuzzleInput("./16_in.txt")
	data := "38006F45291200"

	fmt.Println(data)

	borked := ""

	for _, byte := range data {
		i, err := strconv.ParseUint(string(byte), 16, 4)
    	if err != nil {
        	fmt.Printf("%s", err)
    	}
		fmt.Printf("%04b\n", i)

		borked += fmt.Sprintf("%04b", i)

	}

	fmt.Println(borked, len(borked))

	packets := make([]Packet, 0)

	for {
		newPacket := Packet{}
		newPacket.Version = borked[0:2]
	}

	for i, char := range borked {
		fmt.Println(i, string(char))
		
	}

}

type Packet struct {
	Version uint64
	TypeID uint64
	LengthID uint64
	Subpackets []uint64
} 

func getPuzzleInput(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	
	return bytes
}




