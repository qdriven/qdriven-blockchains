package model

import "gin-vue-admin/global"

type PaletteNode struct {
	global.GVA_MODEL
	NodeAddress    string
	StakeAmount  string
	CurrentBlockNum string
	NodeType int
	SeedNode int
	NodeHost string
	NodeRPCPort string
	NodeP2PPort string
	NodePrivateKey string
	NodeDir string
	Active int
}
