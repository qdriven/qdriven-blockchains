package model

import (
	"gin-vue-admin/global"
)

type NftContract struct {
	global.GVA_MODEL
	TransactionHash         string
	NFTContractAddress string
	Status				uint64
	Name               string
	Symbol             string
	Creator            string
}
