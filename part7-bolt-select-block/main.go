package main

import (
	"log"
	"publicChain/part7-bolt-select-block/blc"

	"github.com/boltdb/bolt"
)

func main() {

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	blockChain := blc.CreateBlockChainWithGenesisBlock(db)

	blockChain.AddBlockToBlockChain("send 100rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash, db)

	blockChain.AddBlockToBlockChain("send 10rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash, db)

	blockChain.AddBlockToBlockChain("send 1rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash, db)

	blockChain.AddBlockToBlockChain("send 1000rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash, db)

	blc.GetAllBlock(db, []byte("l"))
}
