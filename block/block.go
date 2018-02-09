package block

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func New(data []byte, prevHash []byte) *Block {
	b := &Block{
		Timestamp: time.Now().Unix(),
		Data: data,
		PrevBlockHash: prevHash,
	}
	b.SetHash()
	return b
}
