package block

import (
	"math/big"
	"math"
	"bytes"
	"encoding/binary"
	"crypto/sha256"
)

const targetBits = 2 * 8
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))

	return &ProofOfWork{
		target: target,
		block: b,
	}
}

func (pow *ProofOfWork) Run() (int, []byte) {
	nonce := 0
	var hashInt big.Int
	var hash [32]byte

	for nonce < maxNonce {
		data := pow.codifyBlock(nonce)

		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

func (pow *ProofOfWork) codifyBlock(nonce int) []byte {
	nonceInt64 := uint64(nonce)
	var targetBits uint64 = targetBits

	nonceByteSlice := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonceByteSlice, nonceInt64)
	targetBitsByteSlice := make([]byte, 8)
	binary.LittleEndian.PutUint64(targetBitsByteSlice, targetBits)

	data := bytes.Join(
		[][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		nonceByteSlice,
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Validate(nonce int) bool {
	var hashInt big.Int

	data := pow.codifyBlock(nonce)
	hash := sha256.Sum256(data)

	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
