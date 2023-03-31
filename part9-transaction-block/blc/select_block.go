package blc

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func GetNewBlock(db *bolt.DB) (*Block, error) {

	block := &Block{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			return fmt.Errorf("not find table")
		}

		date := b.Get([]byte("l"))

		blockBytes := b.Get(date)

		block = DeserializeBlock(blockBytes)

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
			block.PrintfBlock()
			GetAllBlock(db, block.PrevBlockHash)
		} else {

			if bytes.Equal(prveBlockHash, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) {

				return nil
			}
			date := b.Get(prveBlockHash)
			block = DeserializeBlock(date)
			block.PrintfBlock()
			GetAllBlock(db, block.PrevBlockHash)
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func (block *Block) PrintfBlock() {
	fmt.Printf("Height: %d \n", block.Heigth)
	fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
	fmt.Printf("TimeStamp: %s \n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
	fmt.Printf("Hash: %x \n", block.Hash)
	fmt.Printf("Nonce: %d \n", block.Nonce)
	for _, transcation := range block.Txs {
		fmt.Printf("TxHash: %x \n", transcation.TxHash)
		for _, vins := range transcation.Vins {
			fmt.Printf(" Vins: \n")
			fmt.Printf("    TxHash: %x \n", vins.TxHash)
			fmt.Printf("    ScriptSig: %s \n", vins.ScriptSig)
			fmt.Printf("    Vout: %d \n", vins.Vout)
		}
		for _, vouts := range transcation.Vouts {
			fmt.Printf(" Vouts: \n")
			fmt.Printf("   Value: %d \n", vouts.Value)
			fmt.Printf("   ScriptPubKey: %s \n", vouts.ScriptPubKey)
		}
	}
	fmt.Println("-----------------------")
	fmt.Println()

}
