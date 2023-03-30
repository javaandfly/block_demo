package blc

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func GetNewBlock() (*Block, error) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	block := &Block{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if err != nil || b == nil {
			return fmt.Errorf("create bucket:  %s", err)
		}

		date := b.Get([]byte("l"))

		blockBytes := b.Get(date)

		block = DeserializeBlock(blockBytes)

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block, nil
}

func GetAllBlock(db *bolt.DB, prveBlockHash []byte) error {

	block := &Block{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			return fmt.Errorf("not find Bucket")
		}
		if bytes.Equal([]byte("l"), prveBlockHash) {
			date := b.Get([]byte("l"))
			blockBytes := b.Get(date)
			block = DeserializeBlock(blockBytes)
			block.printfBlock()
			GetAllBlock(db, block.PrevBlockHash)
		} else {

			if bytes.Equal(prveBlockHash, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) {

				return nil
			}
			date := b.Get(prveBlockHash)
			block = DeserializeBlock(date)
			block.printfBlock()
			GetAllBlock(db, block.PrevBlockHash)
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func (block *Block) printfBlock() {
	fmt.Printf("Height: %d \n", block.Heigth)
	fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("TimeStamp: %s \n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
	fmt.Printf("Hash: %x \n", block.Hash)
	fmt.Printf("Nonce: %d \n", block.Nonce)
	fmt.Println("-----------------------")
	fmt.Println()

}
