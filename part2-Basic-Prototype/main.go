package main

import (
	"fmt"
	"publicChain/part2-Basic-Prototype/blc"
)

func main() {

	blockChain := blc.CreateBlockChainWithGenesisBlock()

	blockChain.AddBlockToBlockChain("send 100rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	blockChain.AddBlockToBlockChain("send 10rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	blockChain.AddBlockToBlockChain("send 10rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	blockChain.AddBlockToBlockChain("send 1000rmb to dongzhi", blockChain.Blocks[len(blockChain.Blocks)-1].Heigth+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	fmt.Println(blockChain)

}
