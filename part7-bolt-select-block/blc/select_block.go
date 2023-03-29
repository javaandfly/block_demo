package blc

import (
	"fmt"
	"log"

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

func GetAllBlock(db *bolt.DB) (*Block, error) {

	block := &Block{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			return fmt.Errorf("not find Bucket")
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
