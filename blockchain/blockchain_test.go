package blockchain

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/jgimeno/go-blockchain/block"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestIfItDoesNotHaveGenesisBlockItCreatesIt(t *testing.T) {
	mockedDb := &mockedPersistence{}
	defer mockedDb.AssertExpectations(t)

	mockedDb.On("HasGenesis").Return(false)
	mockedDb.On("Init").Return(nil)
	mockedDb.On("Save", mock.MatchedBy(
		func(b *block.Block) bool {
			return bytes.Equal(b.PrevBlockHash, []byte{})
		},
	)).Return(nil)

	New(mockedDb)
}

func TestItGetsLastHashAsTipWhenItHasGenesis(t *testing.T) {
	mockedDb := &mockedPersistence{}
	defer mockedDb.AssertExpectations(t)

	mockedDb.On("HasGenesis").Return(true)

	lastHash := []byte("TheLastHash")
	mockedDb.On("GetLastHash").Return(lastHash, nil)

	blockchain := New(mockedDb)

	assert.Equal(t, lastHash, blockchain.tip)
}
