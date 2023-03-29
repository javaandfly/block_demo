package blc

import (
	"bytes"
	"encoding/gob"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

func (blcok *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	// block Pack
	err := encoder.Encode(blcok)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(blockBytes []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewBuffer(blockBytes))

	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func (blockChain *BlockChain) AddBlockToBlockChain(date string, height int64, prevHash []byte) {
	newBlock := NewBlock(date, height, prevHash)

	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis block")

	return &BlockChain{[]*Block{genesisBlock}}
}
