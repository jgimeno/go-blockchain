package block_test

import (
	"testing"
	"bytes"
	"github.com/jgimeno/go-blockchain/block"
)

func TestWeCanCreateANewBlock(t *testing.T) {
	g := block.New("The new blog", []byte("previousHash"))

	pow := block.NewProofOfWork(g)

	if !pow.Validate(g.Nonce) {
		t.Fatalf("Error validating created block.")
	}
}

func TestWeCanCreateAGenesisBlock(t *testing.T) {
	g := block.NewGenesis()

	if bytes.Equal(g.Hash, []byte{}) {
		t.Fatal("Error creating genesis block.")
	}
}

func TestWeCanAssertIfABlockIsAGenesisBlock(t *testing.T) {
	notGenesis := block.New("Some data", []byte("PreviusHash"))
	genesis := block.NewGenesis()

	if !genesis.IsGenesis() {
		t.Fatal("Failed asserting genesis block.")
	}

	if notGenesis.IsGenesis() {
		t.Fatal("Failed asserting that noral block is not genesis.")
	}
}

func TestWeCanSerializeAndDeserializeABlock(t *testing.T) {
	b := block.New("Send 200 to Manuel", []byte("TheHash"))

	s := b.Serialize()

	db := block.DeserializeBlock(s)

	if !bytes.Equal(b.Data, db.Data) || !bytes.Equal(b.PrevBlockHash, db.PrevBlockHash) {
		t.Fatalf("Error serializing block.")
	}
}