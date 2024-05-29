package service

import "github.com/ethereum/go-ethereum/common"

func validateAddress(address string) error {
	valid := common.IsHexAddress(address)
	if !valid {
		return InvalidAddressErr
	}
	return nil
}

func validateKeyword(keyword string) error {
	if len(keyword) < 1 {
		return InvalidKeywordErr
	}
	return nil
}
