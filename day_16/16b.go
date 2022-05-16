package main

import (
	"fmt"
	"io/ioutil"
	// "strings"
	"strconv"
)

func identifyPacket(message string, p Packet, versions int) (string, Packet, int) {
	// Take version number and Type ID and return the rest
	// Get version number
	version := message[0:3]
	message = message[3:]
	fmt.Println("version:", version)

	v, err := strconv.ParseInt(version, 2, 8); if err != nil { fmt.Println(err) }
	versions += int(v)

	fmt.Println(versions)

	// Get type ID
	typeID := message[0:3]
	message = message[3:]
	fmt.Println("type ID:", typeID)

	p.TypeID, err = strconv.ParseInt(typeID, 2, 8); if err != nil { fmt.Println(err) }




	return message, p, versions, 

}

func parsePacket(message string, p Packet, version int) (string, Packet, int) {
	message, p, versions, typeID := identifyPacket(message, p, versions)

	switch typeID {
	case 0:

	case 1:

	}
	
}

func main() {
	// data := getPuzzleInput("./16_in.txt")
	// data := "D2FE28" // test 1
	// data := "38006F45291200" // operator, type 1
	data := "EE00D40C823060" // operator, type 0

	message := parseMessage(data)

	fmt.Println(message, len(message))

	packets := make([]Packet, 0)
	p := Packet{Subpackets: make(map[int]int64)}
	
	i := 1
	parsePacket(message, p)

	for {
		if len(message) < 10 {
			break
		}
		fmt.Println("iter", i)
		var err error
		// var tmpInt int64



		
		if newPacket.TypeID == 4 { // Packets with type ID 4 represent a literal value. 
			// Literal value packets encode a single binary number. 
			// To do this, the binary number is padded with leading zeroes until 
			// its length is a multiple of four bits, and then it is broken into groups of four bits. 
			// Each group is prefixed by a 1 bit except the last group, which is prefixed by a 0 bit. 
			// These groups of five bits immediately follow the packet header
			fmt.Println("Case 4 found in iteration", i)
			j := 0
			for {
				// Take next 5
				nextFive := message[0:5]
				fmt.Println("Next five:", nextFive)
				message = message[5:]
				fmt.Println("String left:", message)

				// Prep continue and next 4
				continuing := nextFive[0] - 48
				fmt.Println("Continue? ->", continuing)
				nextFive = nextFive[1:]
				fmt.Println("number:", nextFive)

				// Add all binary to a single string to calculate afterwards
				// tmpInt, err = strconv.ParseInt(nextFive, 2, 8); if err != nil { fmt.Println(err) }
				// str := strconv.Itoa(int(tmpInt))
				// fmt.Println("str:", str)
				newPacket.ValueStr += nextFive

				if continuing == 0 {
					newPacket.ValueInt64, err = strconv.ParseInt(newPacket.ValueStr, 2, len(newPacket.ValueStr))
					
					break
				} else {
					j++
					fmt.Println("case 4, iteration", j)
					continue
				}
			}
		} else {
			fmt.Println("operator packet found in iteration", i)
			// The packet is an "Operator" 
			// Performs some calculations on one or more sub-packets contained within
			// To indicate which subsequent binary data represents its sub-packets,
			// an operator packet can use one of two modes indicated by the bit immediately
			// after the packet header; this is called the length type ID:
				
				
			// Finally, after the length type ID bit and the 15-bit or 11-bit field, the sub-packets appear.

			// Take length ID
			newPacket.LengthID = message[0] - 48
			fmt.Println("length ID:", newPacket.LengthID)
			message = message[1:]

			if newPacket.LengthID == 0 {
				// If the length type ID is 0, then the next 15 bits are a number that 
				// represents the total length in bits of the sub-packets contained by this packet.

				newPacket.SubpacketsArray, err = strconv.ParseInt(message[0:15], 2, 64); if err != nil { fmt.Println(err) }
				message = message[15:]

				for i := 1; i < int(newPacket.SubpacketsArray); i++ {
					fmt.Println("subpacket", i, ":", newPacket.Subpackets[i])
				}
			} else {
				// If the length type ID is 1, then the next 11 bits are a number that 
				// represents the number of sub-packets immediately contained by this packet.

				// Grab the next 11 bits and create new array of subpacket based on the data
				newPacket.SubpacketsArray, err = strconv.ParseInt(message[0:11], 2, 64); if err != nil { fmt.Println(err) }
				fmt.Println(newPacket.SubpacketsArray)
				// subPaArr := make([]int, newPacket.SubpacketsArray)
				message = message[11:]

				// loop through subpackets
				for i := 1; i < int(newPacket.SubpacketsArray); i++ {
					newPacket.Subpackets[i], err = strconv.ParseInt(message[0:11], 0, 64); if err != nil { fmt.Println(err) }
					message = message[11:]
					fmt.Println("subpacket", i, ":", newPacket.Subpackets[i])
				} 
				break

			}

		} 

		packets = append(packets, newPacket)
		fmt.Println(newPacket)
		if len(message) < 3 {
			break
		}

		i++
	}

	// for i, char := range message {
	// 	fmt.Println(i, string(char))
		
	// }

}

type Packet struct {
	Version int64
	TypeID int64
	LengthID byte
	SubpacketsArray int64
	Subpackets map[int]int64
	ValueInt64 int64
	ValueStr string
} 


func parseMessage(data string) string {
	message := ""
	for _, byte := range data {
		i, err := strconv.ParseUint(string(byte), 16, 8)
    	if err != nil {
        	fmt.Printf("%s", err)
    	}
		fmt.Printf("%04b\n", i)

		message += fmt.Sprintf("%04b", i)
	}

	return message
}


func getPuzzleInput(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	
	return bytes
}




