package blockchain

import (
	"testing"
	"bytes"
)

func TestANewBlockChainIncludesAGenesisBlock(t *testing.T) {
	bc := New()

	firstBlock := bc.blocks[0]

	if !firstBlock.IsGenesis() {
		t.Fatal("Failed asserting that the first block of a Blockchain is a Genesis block.")
	}
}

func TestWeCanAddNewBlocks(t *testing.T) {
	bc := New()

	bc.AddBlock("Add 10 btc to Manolo Lama")

	genesisBlock := bc.blocks[0]
	newBlock := bc.blocks[1]

	if !bytes.Equal(newBlock.PrevBlockHash, genesisBlock.Hash) {
		t.Fatal("Failed asserting that new block hash has previous hash from prervius block.")
	}

	if !bytes.Equal(newBlock.Data, []byte("Add 10 btc to Manolo Lama")) {
		t.Fatal("Error asserting expected data for new block.")
	}
}