package blockchain

import "github.com/jgimeno/go-blockchain/block"

type BlockChain struct {
	blocks []*block.Block
}

func (b *BlockChain) AddBlock(data string) {
	prevBlock := b.blocks[len(b.blocks) - 1]
	newBlock := block.New([]byte(data), prevBlock.Hash)
	b.blocks = append(b.blocks, newBlock)
}

func New() *BlockChain {
	return &BlockChain{
		[]*block.Block{block.NewGenesis()},
	}
}
