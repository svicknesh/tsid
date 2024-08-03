package tsid

import (
	"fmt"
	"strings"
)

const (
	timeMask = 18446744073705357312 // the mask to extract time from a number (in hex its '0xffffffffffc00000')
	//randomMask = 4194303              // we need this mask to to get the value of the last 22 bits for random value (in hex it is '0x3fffff')
	strLen = 13
)

// TSIDValue - stores breakdown of TSID
type TSIDValue struct {
	Time, Number    uint64
	NodeID, Counter uint32
	//NodeIDPos uint32
}

// NewFromNumber - creates an instance of `TSIDValue` from a given number
func NewFromNumber(number uint64) (tsv *TSIDValue) {
	tsv = new(TSIDValue)

	if number == 0 {
		return // if there is no valid number, return immediately
	}

	tsv.Number = number

	return
}

// NewFromString - creates an instance of `TSIDValue` from a given string
func NewFromString(str string) (tsv *TSIDValue, err error) {

	len := len(str)

	if len == 0 {
		return nil, fmt.Errorf("newfromstring: no string given") // if there is no valid string, return immediately
	} else if len != strLen {
		return nil, fmt.Errorf("newfromstring: string must be exactly %d characters long, got %d characters ", strLen, len) // the string must be exactly 13 characters long
	}

	tsv = new(TSIDValue)

	bs := uint64(60)
	for _, c := range str {
		///tsv.Number =
		tsv.Number |= alphabets[c] << bs // to maintain the previous values, we need to 'OR' the new value with the existing value
		bs -= 5
	}

	return
}

// Parse - parses the number and crafts the individual components, useful for debugging
func (tsv *TSIDValue) Parse(nodeBit uint64) (err error) {

	if nodeBit > 20 {
		nodeBit = 0 // if the node bitsize goes beyond the boundaries, set it to 0
	}

	counterBit := randomBitSize - nodeBit

	// we split the time from the rest, we don't touch the random part because we don't know the node size yet
	tsv.Time = tsv.Number & timeMask
	//tsv.Counter = uint32(tsv.Number) & randomMask // default assume no node information, everything is counter

	tsv.NodeID = uint32(tsv.Number >> counterBit & ((1 << nodeBit) - 1))
	tsv.Counter = uint32(tsv.Number & ((1 << counterBit) - 1))

	return
}

// String - returns crockord base32 encoding of the number in uppercase
func (tsv *TSIDValue) String() (str string) {

	// fastest way to get the character is to only read the least significant bit of the number and determine which character fits it
	// we shift down 5 bits each time in order to reach the 13 character limit for TSID (12 * 5 bit = 60) and uint64 has a max of 64 bits
	chars := make([]rune, strLen)

	sBit := uint64(60)
	for index := range chars {
		chars[index] = uppercase[tsv.Number>>sBit&31]
		sBit -= 5 //always decrease in 5s
	}

	return string(chars)
}

// LowerCase - returns crockord base32 encoding of the number in lowercase
func (tsv *TSIDValue) LowerCase() (str string) {
	return strings.ToLower(tsv.String())
}
