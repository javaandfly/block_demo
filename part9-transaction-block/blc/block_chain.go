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
type TransactionMessage struct {
	From   []string
	To     []string
	Amount []string
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

func AddBlockToBlockChain(transactionMessage *TransactionMessage, db *bolt.DB) {

	newBlock := &Block{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			return fmt.Errorf("not find table")
		}

		date := b.Get([]byte("l"))

		blockBytes := b.Get(date)

		newBlock = DeserializeBlock(blockBytes)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock = NewBlock(nil, newBlock.Heigth+1, newBlock.Hash)

	err = db.Update(func(tx *bolt.Tx) error {
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

}

func CreateBlockChainWithGenesisBlock(db *bolt.DB, address string) *BlockChain {
	transaction := NewCoinBaseTransaction(address)

	genesisBlock := CreateGenesisBlock([]*Transaction{transaction})

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
