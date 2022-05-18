package main

import "fmt"

func main() {

	bits := []byte{
		// 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0,
		0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1 , 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		// 1,1,1,0,1,1,1,0,0,0,0,0,0,0,0,0,1,1,0,1,0,1,0,0,0,0,0,0,1,1,0,0,1,0,0,0,0,0,1,0,0,0,1,1,0,0,0,0,0,1,1,0,0,0,0,0,
	}

	l, c := readPacket(bits, 0)

	fmt.Println(l, c)
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
	// Get version
	version, count := readBits(data, startpos, 3)
	pos := startpos + count

	// Get typeID
	typeID, count := readBits(data, pos, 3)
	pos += count

	// Get the value of the packet
	switch typeID {
	case 4:
		value, count := readNumber(data, pos)
		pos += count

		return Literal{
			Version: version,
			TypeID: typeID,
			Value: value,
		}, pos - startpos

	default: // Other typeID than 4
		lengthID, count := readBits(data, pos, 1)
		pos += count

		op := Operator{
			Version: version,
			TypeID: typeID,
			LengthID: lengthID,
			Packets: nil,
		}

		if lengthID == 0 {
			length, _ := readBits(data, pos, 15)
			op.Length = length

			subpacketPosStart := pos

			for int64(pos - subpacketPosStart) < op.Length {
				fmt.Printf("Pos: %d, subpacketPosStart: %d, op.length: %d\n", pos, subpacketPosStart, op.Length)
				packet, count := readPacket(data, pos)
				pos += count

				op.Packets = append(op.Packets, packet)
			}

			return op, pos - startpos

			
		} else { // lengthID = 1
			length, count := readBits(data, pos, 11)
			op.Length = length
			// fmt.Println(length)
			pos += count

			op := Operator{
				Version: version,
				TypeID: typeID,
				LengthID: lengthID,
				Length: length,
				Packets: nil,
			}

			for i := int64(0); i < length; i++ {
				// fmt.Println(i, count)

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
		// otherwise, we read 5 more
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