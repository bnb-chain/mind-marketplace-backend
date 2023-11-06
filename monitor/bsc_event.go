package monitor

import (
	"github.com/bnb-chain/mind-marketplace-backend/monitor/contracts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// tx example: https://testnet.bscscan.com/tx/0x60db794687bc0d230bd64fb859df1eb424500c90c801b8bc9857b606781e3e9d#eventlog
const eventBuyTopic = "0xe3d4187f6ca4248660cc0ac8b8056515bac4a8132be2eca31d6d0cc170722a7e"
const eventListTopic = "0x1518ffac149698f404e82117efa8b67d99c365490eefe4e0f91856a93ca8d9c1"
const eventDelistTopic = "0x8fcc1d45240b67aa8f5859c01c295e240be99a9d5e4c11873bb82cf40be7533c"
const eventPriceUpdatedTopic = "0xb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0"

func isTargetEvent(targetTopic ethcommon.Hash, l types.Log) bool {
	return targetTopic.String() == l.Topics[0].String()
}

func parseBuyEvent(l types.Log) (*contracts.MarketplaceBuy, error) {
	if !isTargetEvent(ethcommon.HexToHash(eventBuyTopic), l) {
		return nil, nil
	}

	buy := &contracts.MarketplaceBuy{
		Buyer:   ethcommon.BytesToAddress(l.Topics[1].Bytes()),
		GroupId: big.NewInt(0).SetBytes(l.Topics[2].Bytes()),
	}
	return buy, nil
}

func parseListEvent(l types.Log) (*contracts.MarketplaceList, error) {
	if !isTargetEvent(ethcommon.HexToHash(eventListTopic), l) {
		return nil, nil
	}

	list := &contracts.MarketplaceList{
		Owner:   ethcommon.BytesToAddress(l.Topics[1].Bytes()),
		GroupId: big.NewInt(0).SetBytes(l.Topics[2].Bytes()),
	}
	list.Price = big.NewInt(0).SetBytes(l.Data)
	return list, nil
}

func parseDelistEvent(l types.Log) (*contracts.MarketplaceDelist, error) {
	if !isTargetEvent(ethcommon.HexToHash(eventDelistTopic), l) {
		return nil, nil
	}

	delist := &contracts.MarketplaceDelist{
		Owner:   ethcommon.BytesToAddress(l.Topics[1].Bytes()),
		GroupId: big.NewInt(0).SetBytes(l.Topics[2].Bytes()),
	}
	return delist, nil
}

func parsePriceUpdatedEvent(l types.Log) (*contracts.MarketplacePriceUpdated, error) {
	if !isTargetEvent(ethcommon.HexToHash(eventPriceUpdatedTopic), l) {
		return nil, nil
	}

	updatePrice := &contracts.MarketplacePriceUpdated{
		Owner:   ethcommon.BytesToAddress(l.Topics[1].Bytes()),
		GroupId: big.NewInt(0).SetBytes(l.Topics[2].Bytes()),
	}
	updatePrice.Price = big.NewInt(0).SetBytes(l.Data)
	return updatePrice, nil
}
