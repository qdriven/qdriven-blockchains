package model

import "gin-vue-admin/global"

type ValidatorHistory struct {
	global.GVA_MODEL
	Validator          string
	StackAccount       string
	CurrentBlockHeight int64
	Revoke string
	TxHash string
}

