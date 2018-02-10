package block

import (
	"bytes"
	"time"
	"encoding/gob"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
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

func New(data string, prevHash []byte) *Block {
	b := &Block{
		Timestamp: time.Now().Unix(),
		Data: []byte(data),
		PrevBlockHash: prevHash,
		Hash: []byte{},
		Nonce: 0,
	}

	pw := NewProofOfWork(b)
	n, h := pw.Run()

	b.Hash = h
	b.Nonce = n

	return b
}

func NewGenesis() *Block {
	b := New("Genesis block", []byte{})
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