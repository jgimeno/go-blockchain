package blockchain

import (
	"testing"
)

func TestIfItDoesNotHaveGenesisBlockItCreatesIt(t *testing.T) {
	mockedDb := &mockedDbWithoutGenesis{}
	New(mockedDb)

	if !mockedDb.calledHasGenesis || !mockedDb.calledInit || !mockedDb.calledSave {
		t.Fatalf("Error asserting that blockchain called hasGenesis on creation.")
	}
}
