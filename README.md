# README

A Repo for testing and learning blockchain

## chains-gotest-backend

[chains-gotest-backend](./chains-gotest-backend)
A ERC20/ERC721 restapi for making testing easy for blackchain backend testing:
1. EVM tranasction sender
2. EVM NFT Sender
3. EVM Governance Testing
4. Validators Testing for consensus 
5. Poly Crosschain SmartContract Sender
crosschain examples:

```golang
package sdk

import (
	"chains-gotest-backend/api/evm/evm/binding/plttoken/nftlp"
	"chains-gotest-backend/api/evm/evm/binding/poly/eccd_abi"
	"chains-gotest-backend/api/evm/evm/binding/poly/eccm_abi"
	"chains-gotest-backend/api/evm/evm/binding/poly/eccmp_abi"
	"chains-gotest-backend/api/evm/evm/binding/poly/lock_proxy_abi"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native/plt"
	"github.com/ethereum/go-ethereum/contracts/native/utils"

)

func (c *Client) DeployECCD() (common.Address, error) {
	auth := c.makeDeployAuth()
	addr, tx, _, err := eccd_abi.DeployEthCrossChainData(auth, c.backend)
	if err != nil {
		return utils.EmptyAddress, err
	}
	c.WaitTransaction(tx.Hash())
	return addr, nil
}

func (c *Client) DeployECCM(eccd common.Address, sideChainID uint64) (common.Address, error) {
	auth := c.makeDeployAuth()
	//addr, tx, _, err := eccm_abi.DeployEthCrossChainManager(auth, c.backend, eccd, sideChainID)
	addr, tx, _, err := eccm_abi.DeployEthCrossChainManager(auth, c.backend, eccd)
	if err != nil {
		return utils.EmptyAddress, err
	}
	c.WaitTransaction(tx.Hash())
	return addr, nil
}

func (c *Client) DeployCCMP(eccm common.Address) (common.Address, error) {
	auth := c.makeDeployAuth()
	addr, tx, _, err := eccmp_abi.DeployEthCrossChainManagerProxy(auth, c.backend, eccm)
	if err != nil {
		return utils.EmptyAddress, err
	}
	c.WaitTransaction(tx.Hash())
	return addr, nil
}

func (c *Client) SetPLTCCMP(ccmp common.Address) (common.Hash, error) {
	payload, err := c.packPLT(plt.MethodSetManagerProxy, ccmp)
	if err != nil {
		return utils.EmptyHash, err
	}
	hash, err := c.sendPLTTx(payload)
	if err != nil {
		return utils.EmptyHash, err
	}
	c.WaitTransaction(hash)
	return hash, nil
}

func (c *Client) GetPLTCCMP(blockNum string) (common.Address, error) {
	payload, err := c.packPLT(plt.MethodGetManagerProxy)
	if err != nil {
		return utils.EmptyAddress, err
	}

	enc, err := c.callPLT(payload, blockNum)
	if err != nil {
		return utils.EmptyAddress, err
	}

	var proxy common.Address
	if err := c.unpackPLT(plt.MethodGetManagerProxy, &proxy, enc); err != nil {
		return utils.EmptyAddress, err
	}

	return proxy, nil
}

func (c *Client) BindPLTProxy(targetChainID uint64, targetProxy common.Address) (common.Hash, error) {
	payload, err := c.packPLT(plt.MethodBindProxy, targetChainID, targetProxy.Bytes())
	if err != nil {
		return utils.EmptyHash, err
	}

	hash, err := c.sendPLTTx(payload)
	if err != nil {
		return utils.EmptyHash, err
	}

	c.WaitTransaction(hash)
	return hash, nil
}

func (c *Client) GetBindPLTProxy(targetChainID uint64, blockNum string) (common.Address, error) {
	payload, err := c.packPLT(plt.MethodGetBindedProxy, targetChainID)
	if err != nil {
		return utils.EmptyAddress, err
	}

	enc, err := c.callPLT(payload, blockNum)
	if err != nil {
		return utils.EmptyAddress, err
	}

	var proxy []byte
	if err := c.unpackPLT(plt.MethodGetBindedProxy, &proxy, enc); err != nil {
		return utils.EmptyAddress, err
	}

	return common.BytesToAddress(proxy), nil
}

func (c *Client) BindPLTAsset(targetChainID uint64, targetAsset common.Address) (common.Hash, error) {
	payload, err := c.packPLT(plt.MethodBindAsset, targetChainID, targetAsset.Bytes())
	if err != nil {
		return utils.EmptyHash, err
	}

	hash, err := c.sendPLTTx(payload)
	if err != nil {
		return utils.EmptyHash, err
	}

	c.WaitTransaction(hash)
	return hash, err
}

func (c *Client) GetBindPLTAsset(targetChainID uint64, blockNum string) (common.Address, error) {
	payload, err := c.packPLT(plt.MethodGetBindedAsset, targetChainID)
	if err != nil {
		return utils.EmptyAddress, err
	}

	enc, err := c.callPLT(payload, blockNum)
	if err != nil {
		return utils.EmptyAddress, err
	}

	var asset []byte
	if err := c.unpackPLT(plt.MethodGetBindedAsset, &asset, enc); err != nil {
		return utils.EmptyAddress, err
	}

	return common.BytesToAddress(asset), nil
}

func (c *Client) LockPLT(chainID uint64, dstAddr common.Address, amount *big.Int) (common.Hash, error) {
	payload, err := c.packPLT(plt.MethodLock, chainID, dstAddr.Bytes(), amount)
	if err != nil {
		return utils.EmptyHash, err
	}
	return c.sendPLTTx(payload)
}

func (c *Client) SetNFTCCMP(proxyAddr, ccmp common.Address) (common.Hash, error) {
	proxy, err := lock_proxy_abi.NewLockProxy(proxyAddr, c.backend)
	if err != nil {
		return utils.EmptyHash, err
	}

	auth := c.makeAuth()
	tx, err := proxy.SetManagerProxy(auth, ccmp)
	if err != nil {
		return utils.EmptyHash, err
	}
	c.WaitTransaction(tx.Hash())
	return tx.Hash(), nil
}

func (c *Client) DeployNFTProxy() (common.Address, error) {
	auth := c.makeDeployAuth()
	addr, tx, _, err := nftlp.DeployNFTLockProxy(auth, c.backend)
	if err != nil {
		return utils.EmptyAddress, err
	}
	c.WaitTransaction(tx.Hash())
	return addr, nil
}

func (c *Client) BindNFTProxy(
	localLockProxy common.Address,
	targetLockProxy common.Address,
	targetSideChainID uint64,
) (common.Hash, error) {

	proxy, err := lock_proxy_abi.NewLockProxy(localLockProxy, c.backend)
	if err != nil {
		return utils.EmptyHash, err
	}

	auth := c.makeAuth()
	tx, err := proxy.BindProxyHash(auth, targetSideChainID, targetLockProxy.Bytes())
	if err != nil {
		return utils.EmptyHash, err
	}
	c.WaitTransaction(tx.Hash())
	return tx.Hash(), nil
}

func (c *Client) BindNFTAsset(
	localLockProxy,
	fromAsset,
	toAsset common.Address,
	targetSideChainID uint64,
) (common.Hash, error) {

	proxy, err := nftlp.NewNFTLockProxy(localLockProxy, c.backend)
	if err != nil {
		return utils.EmptyHash, err
	}

	auth := c.makeAuth()
	tx, err := proxy.BindAssetHash(auth, fromAsset, targetSideChainID, toAsset.Bytes())
	if err != nil {
		return utils.EmptyHash, err
	}
	c.WaitTransaction(tx.Hash())
	return tx.Hash(), nil
}

func (c *Client) GetNFTBindAsset(localLockProxy, fromAsset common.Address, targetSideChainID uint64, ) (string, error) {

	proxy, err := nftlp.NewNFTLockProxy(localLockProxy, c.backend)
	if err != nil {
		return "", err
	}

	result, err := proxy.AssetHashMap(nil, fromAsset, targetSideChainID)

	return string(result), err
}

func (c *Client) GetNFTBindProxy(localLockProxy common.Address, targetSideChainID uint64, ) (string, error) {

	proxy, err := nftlp.NewNFTLockProxy(localLockProxy, c.backend)
	if err != nil {
		return "", err
	}
	result, err := proxy.ProxyHashMap(nil, targetSideChainID)
	return string(result), err
}

func (c *Client) InitGenesisBlock(eccmAddr common.Address, rawHdr, publickeys []byte) (common.Hash, error) {
	eccm, err := eccm_abi.NewEthCrossChainManager(eccmAddr, c.backend)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("new EthCrossChainManager err: %s", err)
	}

	auth := c.makeDeployAuth()
	tx, err := eccm.InitGenesisBlock(auth, rawHdr, publickeys)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("call eccm InitGenesisBlock err: %s", err)
	}

	c.WaitTransaction(tx.Hash())
	return tx.Hash(), nil
}

```

## Compound liquidation bot -- demo

- [compound-protocol](compound-protocol)

It is a demo for compound liquidation bot,please do not use in production:

1. Retrieve Compound Protocol API to get healthy indicators
2. If healthy is less than 1, then can liquidate

There are a lot of Ethereum Tools in web3j.


### Others


Several Categories included:

- blockchains: different blockchain networks
- protocols: different protocols in different blockchains
- products: different blockchain/cryptocurrency products on top of different protocols and blockchains
- tech: different tech used in different protocols and blockchains

