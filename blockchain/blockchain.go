package blockchain

import (
	"github.com/jgimeno/go-blockchain/block"
)


type Persistence interface {
	GetLastHash() ([]byte, error)
	Save(*block.Block) error
	HasGenesis() bool
	Init() error
	GetBlockByHash([]byte) *block.Block
}

type BlockChain struct {
	tip []byte
	p Persistence
}

func (bc *BlockChain) Iterator() *Iterator {
	return &Iterator{
		currentHash: bc.tip,
		p: bc.p,
	}
}

func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte

	lastHash, err := bc.p.GetLastHash()

	if err != nil {
		panic("Error getting last hash.")
	}

	newBlock := block.New(data, lastHash)

	err = bc.p.Save(newBlock)

	if err != nil {
		panic("Error saving new block.")
	}

	bc.tip = newBlock.Hash
}

func New(persistence Persistence) *BlockChain {
	var tip []byte

	if !persistence.HasGenesis() {
		persistence.Init()

		genesis := block.NewGenesis()
		persistence.Save(genesis)
		tip = genesis.Hash
	} else {
		lastHash, err := persistence.GetLastHash()
		if err != nil {
			panic("Error creating blockchain.")
		}

		tip = lastHash
	}

	return &BlockChain{
		p:persistence,
		tip:tip,
	}
}

type Iterator struct {
	currentHash []byte
	p Persistence
}

func (bi *Iterator) Next() *block.Block {
	var b *block.Block
	b = bi.p.GetBlockByHash(bi.currentHash)
	bi.currentHash = b.PrevBlockHash
	return b
}
