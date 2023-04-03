package blc

type TXInput struct {
	TxHash []byte

	Vout int64

	ScriptSig string //username
}

func (txInput *TXInput) UnLockWithAddress(address string) bool {

	return txInput.ScriptSig == address
}
