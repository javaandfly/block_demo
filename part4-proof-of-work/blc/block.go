package blc

import (
	"time"
)

type Block struct {
	// Block heigth
	Heigth int64
	// Last Height
	PrevBlockHash []byte
	// Translations date
	Data []byte
	// timestamp
	Timestamp int64
	// hash
	Hash []byte
	// nonce
	Nonce int64
}

func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{Heigth: height, PrevBlockHash: prevBlockHash, Data: []byte(data), Timestamp: time.Now().Unix(), Hash: nil, Nonce: 0}

	// pow or pos
	pow := NewProofOfWork(block)

	//verify
	hash, nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func CreateGenesisBlock(date string) *Block {
	return NewBlock("Genesis block", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
