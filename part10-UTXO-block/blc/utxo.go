package blc

type UTXO struct {
	TXHash []byte

	Index int64

	TXOutput *TXOutput
}
