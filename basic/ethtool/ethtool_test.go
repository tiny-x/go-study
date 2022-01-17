package ethtool

import (
	"fmt"
	"github.com/safchain/ethtool"
	"testing"
)

func Test_ethtool(t *testing.T) {
	ethHandle, err := ethtool.NewEthtool()
	if err != nil {
		panic(err.Error())
	}
	defer ethHandle.Close()

	// Retrieve tx from eth0
	stats, err := ethHandle.Stats("en0")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("TX: %d\n", stats["tx_bytes"])

	// Retrieve peer index of a veth interface
	stats, err = ethHandle.Stats("veth0")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Peer Index: %d\n", stats["peer_ifindex"])
}
