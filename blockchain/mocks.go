package blockchain

import "github.com/jgimeno/go-blockchain/block"

type mockedDbWithoutGenesis struct {
	calledSave bool
	calledHasGenesis bool
	calledInit bool
}

func (*mockedDbWithoutGenesis) GetLastHash() ([]byte, error) {
	panic("implement me")
}

func (m *mockedDbWithoutGenesis) Save(*block.Block) error {
	m.calledSave = true
	return nil
}

func (m *mockedDbWithoutGenesis) HasGenesis() bool {
	m.calledHasGenesis = true
	return false
}

func (m *mockedDbWithoutGenesis) Init() error {
	m.calledInit = true
	return nil
}

func (m *mockedDbWithoutGenesis) CalledSave() bool {
	return m.calledSave
}
