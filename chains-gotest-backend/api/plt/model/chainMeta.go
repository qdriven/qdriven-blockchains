package model

import "gin-vue-admin/global"

type ChainMetaData struct {
	global.GVA_MODEL
	ChainName string
	ChainId   string
	ChainType string //mainnet/testnet
	RPCUrl    string
}

type EvmTransactionRecord struct {
	global.GVA_MODEL
	FromChainId string
	ToChainId   string
	AssetHash   string //mainnet/testnet
	Amount      string
	ToAddress   string
}

type CoinMetaData struct {
	global.GVA_MODEL
	ChainName string
	ChainId   string
	ChainType string //mainnet/testnet
	CoinName  string
	Symbol    string
	Decimal   string
	Hash      string
	Address   string
}

type CrossChainLockProxy struct {
	global.GVA_MODEL
	ChainId          string
	ChainName        string
	ChainType        string
	Active           bool
	LockProxyAddress string
}

type CrossChainAssetMapping struct {
	global.GVA_MODEL
	LockProxyAddress string
	FromAssetHash    string
	ToAssetHash      string
	ToChainId        uint
	Active           bool
}

type CrossChainProxyHashMap struct {
	global.GVA_MODEL
	LockProxyAddress string
	ToProxyHash      string
	ToChainId        uint
}
