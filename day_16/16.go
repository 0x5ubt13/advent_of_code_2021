package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

var Count int
var Value int

func main() {
	bits := parseMessage(getPuzzleInput("16_in.txt"))
	// bits = parseMessage("9C0141080250320F1802104A08")

	packet, _ := readPacket(bits, 0)
	// fmt.Printf("%+v\n", packets)
	partOne := packetArithmeticsVersions(packet)
	partTwo := packetArithmeticsValues(packet)
	// packetArithmeticsValues(packet)

	fmt.Printf("Part 1: -> %d\n", partOne)
	fmt.Printf("Part 2: -> %d\n", partTwo)
}

func packetArithmeticsValues(packet interface{}) int64 {
	// Part 2: calculate values
	switch interf := packet.(type) {
	case Literal:
		return interf.Value
	case Operator:
		// fmt.Println("Type ID =", interf.TypeID)
		switch interf.TypeID {
		case 0: // Sum
			if len(interf.Packets) < 2 {
				return packetArithmeticsValues(interf.Packets[0])
			} else {
				var value int64
				for _, subpacket := range interf.Packets {
					value += packetArithmeticsValues(subpacket)		
				}

				return value
			}
		case 1: // Multiply
			if len(interf.Packets) < 2 {
				return packetArithmeticsValues(interf.Packets[0])
			} else {
				value := 1
				for _, subpacket := range interf.Packets {
					value *= int(packetArithmeticsValues(subpacket))				
				}

				return int64(value)
			}
		case 2: // Min
			if len(interf.Packets) < 2 {
				return packetArithmeticsValues(interf.Packets[0])
			} else {
				var min int64 = 999999999999999
				for _, subpacket := range interf.Packets {
					candidate := packetArithmeticsValues(subpacket)
					if candidate < min {
						min = candidate
					}				
				}

				return min
			}
		case 3: // Max 
			if len(interf.Packets) < 2 {
				return packetArithmeticsValues(interf.Packets[0])
			} else {
				var max int64 
				for _, subpacket := range interf.Packets {
					candidate := packetArithmeticsValues(subpacket)
					if candidate > max {
						max = candidate
					}				
				}

				return max
			}
		case 5: // Greater than (>)
			p1 := packetArithmeticsValues(interf.Packets[0])
			p2 := packetArithmeticsValues(interf.Packets[1])

			if p1 > p2 {
				return 1
			} 

		case 6: // Less than (<)
			p1 := packetArithmeticsValues(interf.Packets[0])
			p2 := packetArithmeticsValues(interf.Packets[1])

			if p1 < p2 {
				return 1
			} 

		case 7: // Equal to (==)
			p1 := packetArithmeticsValues(interf.Packets[0])
			p2 := packetArithmeticsValues(interf.Packets[1])

			if p1 == p2 {
				return 1
			}
		}

	}
	return 0
}

func packetArithmeticsVersions(packet interface{}) int64 {
	var sumVersions int64

	switch interf := packet.(type) {
	case Literal:
		sumVersions += interf.Version
	case Operator:
		sumVersions += interf.Version

		for _, subpacket := range interf.Packets {
			sumVersions += packetArithmeticsVersions(subpacket)
		}
	}

	return sumVersions
}

type Literal struct {
	Version int64
	TypeID int64
	Value int64
} 

type Operator struct {
	Version int64
	TypeID int64
	Value int64
	LengthID int64
	Length int64
	Packets []interface{}
}

// Read packet
func readPacket(data []byte, startpos int) (l interface{}, c int) {
	pos := startpos

	// Get version
	version, count := readBits(data, startpos, 3)
	pos += count
	fmt.Println("Version:", version)

	// Get typeID
	typeID, count := readBits(data, pos, 3)
	pos += count
	fmt.Println("TypeID:", typeID)


	// Get the value of the packet
	switch typeID {
	case 4:
		fmt.Println("Entering typeID = 4 case")
		value, count := readNumber(data, pos)
		pos += count

		return Literal{
			Version: version,
			TypeID: typeID,
			Value: value,
		}, pos - startpos

	default: // Other typeID than 4

		fmt.Println("Entering typeID = not 4 case")

		lengthID, count := readBits(data, pos, 1)
		pos += count
		fmt.Println("length ID =", lengthID)

		op := Operator{
			Version: version,
			TypeID: typeID,
			LengthID: lengthID,
			Packets: nil,
		}

		if lengthID == 0 {
			length, count := readBits(data, pos, 15)
			pos += count
			fmt.Println("Entering length ID = 0. Length =", length)
			
			op.Length = length

			subpacketPosStart := pos

			for int64(pos - subpacketPosStart) < length {
				Count++
				fmt.Println("Entering lengthID=0 loop, iteration number", Count)
				fmt.Printf("Pos: %d, subpacketPosStart: %d, length: %d, op.length: %d\n", pos, subpacketPosStart, length, op.Length)
				
				packet, count := readPacket(data, pos)
				pos += count

				op.Packets = append(op.Packets, packet)
			}

			return op, pos - startpos

		} else { // lengthID = 1
			length, count := readBits(data, pos, 11)
			op.Length = length
			fmt.Println("Entering length ID = 1. Length =", length)

			// fmt.Println(length)
			pos += count
			op.Length = length

			for i := int64(0); i < length; i++ {
				// fmt.Println(i, count)

				fmt.Println("Length ID 1 - iteration number", i)
				packet, count := readPacket(data, pos)
				pos += count

				op.Packets = append(op.Packets, packet)
			}

			return op, pos - startpos
		}
	}
}

func readNumber(data []byte, startpos int) (num int64, count int) {
	for {
		// The part of the last 4 bits is the number we are looking for
		part, _ := readBits(data, startpos, 5)
		num <<= 4
		num |= int64(part & 0x0f)
		count += 5
		startpos += 5

		// if the fifth bit equals 0, finish, that's the last one
		if part & 0x10 == 0 {
			break
		}
		// otherwise, loop starts again -> we read 5 more
	}

	return num, count
}

func readBits(data []byte, startpos, count int) (int64, int) {
	var out int64
	for _, b := range data[startpos:startpos+count] {
	// for i, b := range data[startpos:startpos+count] {
		// fmt.Printf("Reading bits, iter %d: b = %d\n", i, b)
		// fmt.Printf("Current out:\nDecimal: %d\t Binary: %b\n", out, out)
		out <<= 1
		// fmt.Printf("New out after bitshifting left 1:\nDecimal: %d\t Binary: %b\n", out, out)
		out |= int64(b)
		// fmt.Printf("New out after out or-equals int64(b):\nDecimal: %d\t Binary: %b\n\n\n", out, out)
	}

	return out, count
}

func parseMessage(data string) []byte {
	// fmt.Println(data)
	message := ""
	bits := make([]byte, 0)
	for _, byte := range string(data) {
		i, err := strconv.ParseUint(string(byte), 16, 8)
    	if err != nil {
        	fmt.Printf("%s", err)
    	}
		// fmt.Printf("%04b\n", i)

		message += fmt.Sprintf("%04b", i)
	}

	for _, bit := range message {
		newBit, err := strconv.ParseUint(string(bit), 16, 8); if err != nil { panic(err) }
		bits = append(bits, byte(newBit))
	}

	return bits
}

func getPuzzleInput(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(bytes)

	message := ""

	for _, num := range string(bytes) {
		message += string(num)
	}

	// fmt.Println(message)
	
	return message
}