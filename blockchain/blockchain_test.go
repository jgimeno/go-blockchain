package blockchain

import (
	"testing"
)

func TestANewBlockChainIncludesAGenesisBlock(t *testing.T) {
	bc := New()

	firstBlock := bc.blocks[0]

	if !firstBlock.IsGenesis() {
		t.Fatal("Failed asserting that the first block of a Blockchain is a Genesis block.")
	}
}
