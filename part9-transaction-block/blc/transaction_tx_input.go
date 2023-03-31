package blc

type TXInput struct {
	TxHash []byte

	Vout int64

	ScriptSig string //username
}

func NewTXInput() {

}
