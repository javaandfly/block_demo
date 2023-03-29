package blc

type BlockChain struct {
	Blocks []*Block
}

func (blockChain *BlockChain) AddBlockToBlockChain(date string, height int64, prevHash []byte) {
	newBlock := NewBlock(date, height, prevHash)

	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis block")

	return &BlockChain{[]*Block{genesisBlock}}
}
