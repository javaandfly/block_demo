package blc

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
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

func (blockChain *BlockChain) AddBlockToBlockChain(date string, height int64, prevHash []byte, db *bolt.DB) {
	newBlock := NewBlock(date, height, prevHash)

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			return fmt.Errorf("not find table")
		}

		blockBytes := newBlock.Serialize()
		err := b.Put([]byte(newBlock.Hash), blockBytes)
		if err != nil {
			return err
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

func CreateBlockChainWithGenesisBlock(db *bolt.DB) *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis block")

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("BlockBucket"))
		if err != nil || b == nil {
			return fmt.Errorf("create bucket:  %s", err)
		}
		blockBytes := genesisBlock.Serialize()
		err = b.Put([]byte(genesisBlock.Hash), blockBytes)
		if err != nil {
			return err
		}

		err = b.Put([]byte("l"), genesisBlock.Hash)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return &BlockChain{[]*Block{genesisBlock}}
}
