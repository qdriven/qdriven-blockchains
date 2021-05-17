package model

import "gin-vue-admin/global"

type CrossChainTx struct {
	global.GVA_MODEL
	FromChain   string
	FromAddress string
	ToAddress   string
	Amount      string
	TxHash      string
	ToChain     string
	Status	string
}
