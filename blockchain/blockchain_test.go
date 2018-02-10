package blockchain

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/jgimeno/go-blockchain/block"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/jgimeno/go-blockchain/blockchain/mocks"
)

func TestIfItDoesNotHaveGenesisBlockItCreatesIt(t *testing.T) {
	mockedDb := &mocks.Persistence{}
	defer mockedDb.AssertExpectations(t)

	mockedDb.On("HasGenesis").Return(false)
	mockedDb.On("Init").Return(nil)
	mockedDb.On("Save", mock.MatchedBy(
		func(b *block.Block) bool {
			return bytes.Equal(b.PrevBlockHash, []byte{})
		},
	)).Return(nil)

	p := New(mockedDb)

	t.Run("We add a block correctly.", func(t *testing.T) {
		mockedDb.On("GetLastHash").Return([]byte("theLastHash"), nil)
		mockedDb.On("Save", mock.MatchedBy(
			func(b *block.Block) bool {
				return bytes.Equal(b.PrevBlockHash, []byte("theLastHash"))
			},
		)).Return(nil)


		p.AddBlock("New block.")
	})

	t.Run("We can get an iterator", func(t *testing.T) {
		i := p.Iterator()

		if i.p != mockedDb {
			t.Fatalf("Error checking that iterator has an instance of persistence db.")
		}

		if !bytes.Equal(i.currentHash, p.tip) {
			t.Fatalf("Error checking that the iterator has the lastHash same as the tip of blockchain.")
		}
	})
}

func TestItGetsLastHashAsTipWhenItHasGenesis(t *testing.T) {
	mockedDb := &mocks.Persistence{}
	defer mockedDb.AssertExpectations(t)

	mockedDb.On("HasGenesis").Return(true)

	lastHash := []byte("TheLastHash")
	mockedDb.On("GetLastHash").Return(lastHash, nil)

	blockchain := New(mockedDb)

	assert.Equal(t, lastHash, blockchain.tip)
}

func TestIterator(t *testing.T) {
	mockedDb := &mocks.Persistence{}
	defer mockedDb.AssertExpectations(t)

	i := &Iterator{
		currentHash: []byte("TheCurrentHash"),
		p: mockedDb,
	}

	t.Run(
		"When I get next block it changes the next hash as the prevhash of the returned block",
		func(t *testing.T) {
			assert.Equal(t, []byte("TheCurrentHash"), i.currentHash)

			b := block.New("previous block", []byte("prevHash"))

			mockedDb.On("GetBlockByHash", []byte("TheCurrentHash")).Return(b)
			nextBlock := i.Next()

			assert.Equal(t, nextBlock.PrevBlockHash, i.currentHash)
		},
	)
}
