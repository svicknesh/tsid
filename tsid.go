package tsid

import (
	"math/rand/v2"
	"time"
)

const (
	epoch         = 1577808000000 // 2020-01-01 00:00:00 UTC in miliseconds
	randomBitSize = 22            // combined node and counter will always equal 22
)

// TSID - Time-Sorted Unique Identifiers structure
type TSID struct {
	node *Node
}

// New - creates new instance of TSID with customizable `Option`
func New(node *Node) (t *TSID) {
	t = new(TSID)
	t.node = node
	return
}

// NewDefault - creates new instance of TSID with library specific defaults
func NewDefault() (t *TSID) {
	t = new(TSID)
	t.node, _ = NewNode(0, 0) // default has no node size
	return
}

// Generate - generates a TSID
func (t *TSID) Generate() (tsidValue *TSIDValue) {

	tsidValue = new(TSIDValue)

	// we only want the first 42 bits of the time
	mili := time.Now().UnixMilli()
	tsidValue.Time = (uint64(mili) - epoch) << randomBitSize

	counter := rand.Uint64() & t.node.counterBits // using 'math/rand/v2'
	//counter := uint64(rand.NewSource(mili).Int63()) & t.node.counterBits // using 'math/rand'

	random := t.node.idPos | counter // combine the node id int he expected position with the randomly generated counter to get the final random value

	tsidValue.Number = tsidValue.Time | random

	//tsidValue.NodeIDPos = uint32(t.node.idPos)
	//tsidValue.NodeID = uint32(t.node.id)
	//tsidValue.Counter = uint32(counter)

	return

}
