package block

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
	"encoding/gob"
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

func (b *Block) IsGenesis() bool {
	return bytes.Equal(b.PrevBlockHash, []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		panic("Error encoding block.")
	}

	return result.Bytes()
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

func NewGenesis() *Block {
	b := New([]byte("Genesis block"), []byte{})
	return b
}

func DeserializeBlock(encodedBlock []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(encodedBlock))
	err := decoder.Decode(&block)

	if err != nil {
		panic("Error decoding block.")
	}

	return &block
}