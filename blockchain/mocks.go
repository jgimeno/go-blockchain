package blockchain

import (
	"github.com/jgimeno/go-blockchain/block"
	"github.com/stretchr/testify/mock"
)

type mockedPersistence struct {
	mock.Mock
}

func (m *mockedPersistence) GetLastHash() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (m *mockedPersistence) Save(b *block.Block) error {
	args := m.Called(b)
	return args.Error(0)
}

func (m *mockedPersistence) HasGenesis() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *mockedPersistence) Init() error {
	args := m.Called()
	return args.Error(0)
}
