package tsid

import (
	"fmt"
)

// Node - details about the node
type Node struct {
	idPos, id   uint64
	counterBits uint64
}

func NewNode(bitsize uint64, nodeid uint64) (node *Node, err error) {
	node = new(Node)

	if bitsize > 20 {
		//bitsize = 0 // if the node bitsize goes beyond the boundaries, set it to 0
		return nil, fmt.Errorf("newnode: node bitsize is between 0-20, got %d", bitsize)
	}

	// make sure the node id fits into the bitsize
	nodeMax := uint64(1 << bitsize)
	if nodeid >= nodeMax {
		return nil, fmt.Errorf("newnode: node id exceeds maximum bitsize value, given %d, maximum %d", nodeid, nodeMax-1)
	}

	bsCounter := randomBitSize - bitsize // this is the bit size available for counter

	// get the node value and shift the bits by counter size, which will give us the position to put the node id
	node.idPos = uint64(nodeid << bsCounter) // craft the node id in the expected position
	node.id = nodeid

	node.counterBits = (1 << bsCounter) - 1 // this will give us all 1's for the counter

	return
}
