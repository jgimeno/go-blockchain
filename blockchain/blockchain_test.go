package blockchain

import (
	"testing"
	"fmt"
)

func TestANewBlockChainIncludesAGenesisBlock(t *testing.T) {
	bc := New()
	fmt.Printf("%s", string(bc.tip))
}
