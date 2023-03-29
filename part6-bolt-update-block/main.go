package main

import (
	"fmt"
	"log"
	"publicChain/part6-bolt-update-block/blc"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	block := &blc.Block{}
	blockChain := blc.CreateBlockChainWithGenesisBlock(db)
	defer db.Close()

	blockChain.AddBlockToBlockChain("send 100rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash, db)
	block, err = blc.GetNewBlock(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block)
	blockChain.AddBlockToBlockChain("send 10rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash, db)
	block, err = blc.GetNewBlock(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block)

	blockChain.AddBlockToBlockChain("send 1rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash, db)
	block, err = blc.GetNewBlock(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block)

	blockChain.AddBlockToBlockChain("send 1000rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash, db)

	block, err = blc.GetNewBlock(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block)
}
