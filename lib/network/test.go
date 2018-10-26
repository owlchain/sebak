//
// Provides test utility to mock a network in unittests
//
package network

import (
	"boscoin.io/sebak/lib/node"

	"github.com/stellar/go/keypair"
)

//
// Create a MemoryNetwork for unittests purpose
//
// If the argument `prev` is `nil`, a whole new network, disconnected from any other network,
// is created. If `prev` is not `nil`, the returned `MemoryNetwork` will be reachable from
// every node `prev` can reference
//
func CreateMemoryNetwork(prev *MemoryNetwork) (*keypair.Full, *MemoryNetwork, *node.LocalNode) {
	mn := prev.NewMemoryNetwork()

	kp, _ := keypair.Random()
	localNode, _ := node.NewLocalNode(kp, mn.Endpoint(), "")

	mn.SetLocalNode(localNode)

	return kp, mn, localNode
}
