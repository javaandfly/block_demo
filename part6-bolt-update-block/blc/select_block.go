package blc

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func GetNewBlock(db *bolt.DB) (*Block, error) {

	block := &Block{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			return fmt.Errorf("get BlockBucket error")
		}

		date := b.Get([]byte("l"))

		blockBytes := b.Get(date)

		fmt.Println(blockBytes)

		block = DeserializeBlock(blockBytes)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block, nil
}
