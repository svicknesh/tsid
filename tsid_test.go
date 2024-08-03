package tsid_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/svicknesh/tsid"
)

func TestTSIDGen(t *testing.T) {

	//node, err := tsid.NewNode(10, 1023)
	node, err := tsid.NewNode(6, 63)
	if nil != err {
		log.Println(err)
		os.Exit(1)
	}

	ts := tsid.New(node)

	//ts := tsid.NewDefault()

	tsidValue := ts.Generate()
	fmt.Println(tsidValue.Number)

	//tsidValue.Number = 432511671823499267
	fmt.Println(tsidValue.String())

	//tsidValue.Parse(6)
	//fmt.Println(tsidValue.Number, tsidValue.Time, tsidValue.NodeID, tsidValue.Counter)

}

func TestTSIDStr(t *testing.T) {

	// sample number '432511671823499267' with string '0C04Q2BR40003' for verification
	// bitmask for above number is '0000011000000000100101110001001011110000010000000000000000000011'

	// use the string to generate the number again
	tsidValue2, err := tsid.NewFromString("0C04Q2BR40003")
	if nil != err {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(tsidValue2.Number, tsidValue2.String())

}

func TestTSIDNum(t *testing.T) {

	// sample number '432511671823499267' with string '0C04Q2BR40003' for verification
	// bitmask for above number is '0000011000000000100101110001001011110000010000000000000000000011'

	// use the string to generate the number again
	tsidValue2 := tsid.NewFromNumber(432511671823499267)
	fmt.Println(tsidValue2)

}

func TestTSIDSplit(t *testing.T) {

	nodeLen := uint64(8)

	//node, err := tsid.NewNode(10, 1023)
	node, err := tsid.NewNode(nodeLen, 254)
	if nil != err {
		log.Println(err)
		os.Exit(1)
	}

	ts := tsid.New(node)

	//ts := tsid.NewDefault()

	tsidValue := ts.Generate()
	fmt.Println(tsidValue.Number)

	tsidValue2, err := tsid.NewFromString(tsidValue.String())
	if nil != err {
		log.Println(err)
		os.Exit(1)
	}

	tsidValue2.Parse(nodeLen)
	fmt.Println(tsidValue2.Number, tsidValue2.Time, tsidValue2.NodeID, tsidValue2.Counter)

}

func BenchmarkTSID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		node, err := tsid.NewNode(10, 1023)
		if nil != err {
			log.Println(err)
			os.Exit(1)
		}
		_ = node
	}
}
