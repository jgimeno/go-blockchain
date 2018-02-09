package block_test

import (
	"testing"
	"crypto/sha256"
	"strconv"
	"bytes"
	"github.com/jgimeno/go-blockchain/block"
)

func TestWeCanCreateABlock(t *testing.T) {
	prevHash := sha256.Sum256([]byte("Previus Hash"))

	b := block.New([]byte("Send 200 to Manuel"), prevHash[:])

	if !bytes.Equal(b.Hash, generateHash(b.Timestamp, prevHash[:])) {
		t.Fatalf("Error generating the block hash.")
	}
}

func generateHash(timestamp int64, prevHash []byte) []byte {
	data := []byte("Send 200 to Manuel")
	t := []byte(strconv.FormatInt(timestamp, 10))
	tSlice := t[:]

	headers := bytes.Join([][]byte{prevHash, data, tSlice}, []byte{})

	hash := sha256.Sum256(headers)
	return hash[:]
}

func TestWeCanCreateAGenesisBlock(t *testing.T) {
	g := block.NewGenesis()

	if bytes.Equal(g.Hash, []byte{}) {
		t.Fatal("Error creating genesis block.")
	}
}

func TestWeCanAssertIfABlockIsAGenesisBlock(t *testing.T) {
	notGenesis := block.New([]byte("Some data"), []byte("PreviusHash"))
	genesis := block.NewGenesis()

	if !genesis.IsGenesis() {
		t.Fatal("Failed asserting genesis block.")
	}

	if notGenesis.IsGenesis() {
		t.Fatal("Failed asserting that noral block is not genesis.")
	}
}

func TestWeCanSerializeAndDeserializeABlock(t *testing.T) {
	b := block.New([]byte("Send 200 to Manuel"), []byte("TheHash"))

	s := b.Serialize()

	db := block.DeserializeBlock(s)

	if !bytes.Equal(b.Data, db.Data) || !bytes.Equal(b.PrevBlockHash, db.PrevBlockHash) {
		t.Fatalf("Error serializing block.")
	}
}