package blc

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// UTXO
type Transaction struct {
	TxHash []byte

	Vins []*TXInput

	//user
	Vouts []*TXOutput
}

func (tx *Transaction) IsCionbaseTransaction() bool {
	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1
}

func (transaction *Transaction) HashTransaction() {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(transaction)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())
	transaction.TxHash = hash[:]

}

// create genesis block
func NewCoinBaseTransaction(address string) *Transaction {

	//Consumption
	txInput := &TXInput{[]byte{}, -1, "Genesis Block"}

	txOutput := &TXOutput{10, address}

	txCoinBase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOutput}}

	txCoinBase.HashTransaction()

	return txCoinBase
}

func NewTransaction(transactionMessage *TransactionMessage) *Transaction {
	txInputs := []*TXInput{}
	txOutputs := []*TXOutput{}
	for i := 0; i < len(transactionMessage.From); i++ {
		txInputs[i] = &TXInput{[]byte{}, int64(i), transactionMessage.From[0]}
		txOutputs[i] = &TXOutput{}
	}
	return nil
}
