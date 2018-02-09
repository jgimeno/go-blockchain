package pow

import (
	"testing"
	"github.com/jgimeno/go-blockchain/block"
)

func TestUsageOfProofOfWork(t *testing.T) {
	b := block.New([]byte("The data"), []byte("PrevHash"))
	pow := New(b)

	_, h := pow.Run()

	numBytes := 2

	for i := 0; i < numBytes; i++ {
		if h[i] != 0 {
			t.Fatalf("Failed asserting that byte num %v is 0, %v instead.")
		}
	}
}

func TestValidateProofOfWork(t *testing.T) {
	b := block.New([]byte("The data"), []byte("PrevHash"))
	pow := New(b)

	n, _ := pow.Run()

	if !pow.Validate(n) {
		t.Fatalf("Error validating proof of work.")
	}
}


