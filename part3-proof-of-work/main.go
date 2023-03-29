package main

import (
	"fmt"
	"publicChain/part3-proof-of-work/blc"
)

func main() {

	block := blc.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Println(block.Hash)
	fmt.Println(block.Nonce)

	proofOfWork := blc.NewProofOfWork(block)

	fmt.Println(proofOfWork.IsVaild())
}
