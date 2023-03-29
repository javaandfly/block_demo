package blc

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

func (block *Block) SetHash() {

	heightBytes := IntToHex(block.Heigth)

	timeString := strconv.FormatInt(block.Timestamp, 2)

	timeBytes := []byte(timeString)

	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})

	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{Heigth: height, PrevBlockHash: prevBlockHash, Data: []byte(data), Timestamp: time.Now().Unix(), Hash: nil}
	block.SetHash()
	return block
}

func CreateGenesisBlock(date string) *Block {
	return NewBlock("Genesis block", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
