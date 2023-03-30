package blc

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//hash 0000 0000 0000 0000 1001 0101 1101 ... 1101, 256bit

// I expect top 16 bits is 0
const targetBit = 10

type ProofOfWork struct {
	Block *Block //Currently validated blocks

	//1,create Initial value = 1
	//2,<-- 256 - targetBit = xxxx if < this Target is ok
	Target *big.Int // Big data storage,Data difficulty
}

func (pow *ProofOfWork) IsVaild() bool {
	var hashInt big.Int
	hashInt.SetBytes(pow.Block.Hash)

	return pow.Target.Cmp(&hashInt) == 1
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(pow.Block.Heigth)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
	nonce := 0

	var hashInt big.Int

	var hash [32]byte
	for {
		dataBytes := pow.prepareData(nonce)
		//create hash
		hash = sha256.Sum256(dataBytes)

		fmt.Printf("\r%x", hash)
		//storage hash with hashInt
		hashInt.SetBytes(hash[:])
		//judge hashInt and tatget
		if pow.Target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++
	}
	fmt.Println()
	return hash[:], int64(nonce)
}

// build pow class
func NewProofOfWork(block *Block) *ProofOfWork {

	target := big.NewInt(1)

	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}
