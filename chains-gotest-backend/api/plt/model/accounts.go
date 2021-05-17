package model

import (
	"gin-vue-admin/global"
)

type TestAccount struct {
	global.GVA_MODEL
	Address        string
	AccIndex       int
	DerivationPath string
	PrivateKey     string
	IsAdmin        bool
}

type TestAccountBalance struct {
	global.GVA_MODEL
	Address         string
	Balance         uint64
	CurrentBlockNum int
	Symbol string
	ChainId int
}
