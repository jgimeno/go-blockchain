package block_test

import (
	"testing"
	"github.com/jgimeno/blockchain/block"
	"crypto/sha256"
	"strconv"
	"bytes"
)

func TestWeCanCreateABlock(t *testing.T) {
	prevHash := sha256.Sum256([]byte("Previus Hash"))

	b := block.NewBlock([]byte("Send 200 to Manuel"), prevHash[:])

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