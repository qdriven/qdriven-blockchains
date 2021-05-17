package model

import "gin-vue-admin/global"

type GlobalParameter struct {
	global.GVA_MODEL
	CurrentBlockNo    string
	NftGasPriceInPlt  string
	NftMintPriceInPlt string
	RewardPeriod string
}