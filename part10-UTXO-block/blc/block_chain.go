package blc

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type BlockChain struct {
	Blocks []*Block
}
type TransactionMessage struct {
	From   []string
	To     []string
	Amount []string
}

func (blcok *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	// block Pack
	err := encoder.Encode(blcok)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(blockBytes []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewBuffer(blockBytes))

	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func AddBlockToBlockChain(transactionMessage *TransactionMessage, db *bolt.DB) {

	newBlock := &Block{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			return fmt.Errorf("not find table")
		}

		date := b.Get([]byte("l"))

		blockBytes := b.Get(date)

		newBlock = DeserializeBlock(blockBytes)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock = NewBlock(nil, newBlock.Heigth+1, newBlock.Hash)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			return fmt.Errorf("not find table")
		}

		blockBytes := newBlock.Serialize()
		err := b.Put([]byte(newBlock.Hash), blockBytes)
		if err != nil {
			return err
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

func CreateBlockChainWithGenesisBlock(db *bolt.DB, address string) *BlockChain {
	transaction := NewCoinBaseTransaction(address)

	genesisBlock := CreateGenesisBlock([]*Transaction{transaction})

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("BlockBucket"))
		if err != nil || b == nil {
			return fmt.Errorf("create bucket:  %s", err)
		}
		blockBytes := genesisBlock.Serialize()
		err = b.Put([]byte(genesisBlock.Hash), blockBytes)
		if err != nil {
			return err
		}

		err = b.Put([]byte("l"), genesisBlock.Hash)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return &BlockChain{[]*Block{genesisBlock}}
}

func GetUTXOWithAddress(db *bolt.DB, prveBlockHash []byte, address string, inMap map[string][]int64, utxos []*UTXO) error {
	block := &Block{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			return fmt.Errorf("not find Bucket")
		}
		if bytes.Equal([]byte("l"), prveBlockHash) {
			date := b.Get([]byte("l"))
			blockBytes := b.Get(date)
			block = DeserializeBlock(blockBytes)
			block.PrintfBlock()
			for _, tx := range block.Txs {
				for _, in := range tx.Vins {
					if in.UnLockWithAddress(address) {
						key := hex.EncodeToString(in.TxHash)

						inMap[key] = append(inMap[key], in.Vout)
					}
				}

				for index, out := range tx.Vouts {
					if out.UnLockScriptPubKeyWithAddress(address) {
						if len(inMap) != 0 {
							for txHash, indexArrat := range inMap {
								if txHash == hex.EncodeToString(tx.TxHash) {
									for _, i := range indexArrat {
										if int64(index) == i {
											continue
										} else {
											utxos = append(utxos, &UTXO{TXHash: tx.TxHash, Index: int64(index), TXOutput: out})
										}
									}
								} else {
									utxos = append(utxos, &UTXO{TXHash: tx.TxHash, Index: int64(index), TXOutput: out})
								}
							}
						} else {
							utxos = append(utxos, &UTXO{TXHash: tx.TxHash, Index: int64(index), TXOutput: out})
						}
					}
				}
			}
			GetUTXOWithAddress(db, block.PrevBlockHash, address, inMap, utxos)
		} else {

			if bytes.Equal(prveBlockHash, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) {

				return nil
			}

			date := b.Get(prveBlockHash)
			block = DeserializeBlock(date)
			block.PrintfBlock()
			for _, tx := range block.Txs {
				for _, in := range tx.Vins {
					if in.UnLockWithAddress(address) {
						key := hex.EncodeToString(in.TxHash)

						inMap[key] = append(inMap[key], in.Vout)
					}
				}

				for index, out := range tx.Vouts {
					if out.UnLockScriptPubKeyWithAddress(address) {
						if len(inMap) != 0 {
							for txHash, indexArrat := range inMap {
								if txHash == hex.EncodeToString(tx.TxHash) {
									for _, i := range indexArrat {
										if int64(index) == i {
											continue
										} else {
											utxos = append(utxos, &UTXO{TXHash: tx.TxHash, Index: int64(index), TXOutput: out})
										}
									}
								} else {
									utxos = append(utxos, &UTXO{TXHash: tx.TxHash, Index: int64(index), TXOutput: out})
								}
							}
						} else {
							utxos = append(utxos, &UTXO{TXHash: tx.TxHash, Index: int64(index), TXOutput: out})
						}
					}
				}
			}
			GetUTXOWithAddress(db, block.PrevBlockHash, address, inMap, utxos)
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func GetBanlance(db *bolt.DB, address string, prveBlockHash []byte) int64 {
	inMap := make(map[string][]int64)
	UTXOs := []*UTXO{}
	err := GetUTXOWithAddress(db, prveBlockHash, address, inMap, UTXOs)
	if err != nil {
		log.Panic(err)
	}
	var money int64
	for _, UTXO := range UTXOs {
		money = money + UTXO.TXOutput.Value
	}
	return money
}
