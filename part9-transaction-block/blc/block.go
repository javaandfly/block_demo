package blc

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	// Block heigth
	Heigth int64
	// Last Height
	PrevBlockHash []byte
	// Translations date
	Txs []*Transaction
	// timestamp
	Timestamp int64
	// hash
	Hash []byte
	// nonce
	Nonce int64
}

func (block *Block) HashTransaction() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

func NewBlock(data []*Transaction, height int64, prevBlockHash []byte) *Block {
	block := &Block{Heigth: height, PrevBlockHash: prevBlockHash, Txs: data, Timestamp: time.Now().Unix(), Hash: nil, Nonce: 0}

	// pow or pos
	pow := NewProofOfWork(block)

	//verify
	hash, nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func CreateGenesisBlock(date []*Transaction) *Block {
	return NewBlock(date, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
