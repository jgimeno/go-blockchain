package boltdb_test

import (
	"testing"
	"github.com/jgimeno/go-blockchain/boltdb"
	"os"
	"github.com/jgimeno/go-blockchain/block"
	"bytes"
)

func TestItSavesABlock(t *testing.T) {
	p := boltdb.New()
	defer os.Remove("culo")

	if p.HasGenesis() {
		t.Fatal("Failed to check that it does not have genesis on first instantiation.")
	}

	p.Init()

	if !p.HasGenesis() {
		t.Fatalf("Failed asserting that after init it has genesis.")
	}

	b := block.New("New data", []byte("ThePrevHash"))
	p.Save(b)

	lastHash, _ := p.GetLastHash()
	if !bytes.Equal(lastHash, b.Hash) {
		t.Fatalf("Failed asserting that last hash in db %x is the same as hash of new block. %x", lastHash, b.Hash)
	}
}
