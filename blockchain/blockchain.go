package blockchain

import (
	"github.com/coreos/bbolt"
	"github.com/jgimeno/go-blockchain/block"
)

const blockBucket = "Block"

type BlockChain struct {
	tip []byte
	db *bolt.DB
}

func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		panic("Cannot get last block.")
	}

	newBlock := block.New(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			return  err
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			return  err
		}

		bc.tip = newBlock.Hash

		return nil
	})

	if err != nil {
		panic("Error adding block to the block chain.")
	}
}

func New() *BlockChain {
	var tip []byte

	db, err := bolt.Open("culo", 0600, nil)
	if err != nil {
		panic("Error while opening the file.")
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			genesis := block.NewGenesis()

			b, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				return err
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				return err
			}

			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	return &BlockChain{
		tip, db,
	}
}
