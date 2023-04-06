package blc

type TXOutput struct {
	Value        int64
	ScriptPubKey string //pub_pra
}

func (txOutput *TXOutput) UnLockScriptPubKeyWithAddress(address string) bool {

	return txOutput.ScriptPubKey == address
}
