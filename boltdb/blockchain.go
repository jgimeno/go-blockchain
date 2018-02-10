package boltdb

import (
	"github.com/coreos/bbolt"
	"github.com/jgimeno/go-blockchain/block"
)

const blockBucket = "Block"

func New() *Persistence {
	p, _ := bolt.Open("culo", 0600, nil)
	return &Persistence{p}
}

type Persistence struct {
	b *bolt.DB
}

func (p *Persistence) HasGenesis() bool {
	var b *bolt.Bucket
	p.b.View(func(tx *bolt.Tx) error {
		b = tx.Bucket([]byte(blockBucket))
		return nil
	})

	if b == nil {
		return false
	} else {
		return true
	}
}

func (p *Persistence) Init() error {
	p.b.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(blockBucket))
		if err != nil {
			return err
		}
		return nil
	})

	return nil
}

func (p *Persistence) GetLastHash() ([]byte, error) {
	var lastHash []byte

	err := p.b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		return nil, err
	}

	return lastHash, nil
}

func (p *Persistence) Save(block *block.Block) error {
	err := p.b.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		err := b.Put(block.Hash, block.Serialize())

		if err != nil {
			return  err
		}

		err = b.Put([]byte("l"), block.Hash)
		if err != nil {
			return  err
		}
		return nil
	})

	return err
}

