// Code generated by mockery v1.0.0
package mocks

import block "github.com/jgimeno/go-blockchain/block"

import mock "github.com/stretchr/testify/mock"

// Persistence is an autogenerated mock type for the Persistence type
type Persistence struct {
	mock.Mock
}

// GetBlockByHash provides a mock function with given fields: _a0
func (_m *Persistence) GetBlockByHash(_a0 []byte) *block.Block {
	ret := _m.Called(_a0)

	var r0 *block.Block
	if rf, ok := ret.Get(0).(func([]byte) *block.Block); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*block.Block)
		}
	}

	return r0
}

// GetLastHash provides a mock function with given fields:
func (_m *Persistence) GetLastHash() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasGenesis provides a mock function with given fields:
func (_m *Persistence) HasGenesis() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Init provides a mock function with given fields:
func (_m *Persistence) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: _a0
func (_m *Persistence) Save(_a0 *block.Block) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*block.Block) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
