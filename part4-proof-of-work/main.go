package main

import (
	"fmt"
	"publicChain/part4-proof-of-work/blc"
)

func main() {

	block := blc.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Println(block.Hash)
	fmt.Println(block.Nonce)

	bytes := block.Serialize()
	blockTow := blc.DeserializeBlock(bytes)

	fmt.Println(blockTow)
}
