package blc

type BlockChain struct {
	Blocks []*Block
}

func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis block")

	return &BlockChain{[]*Block{genesisBlock}}
}
