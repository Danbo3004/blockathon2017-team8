package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func NonSignerFunc(s types.Signer, a common.Address, t *types.Transaction) (*types.Transaction, error) {
	return t, nil
}

func to32Byte(str string) [32]byte {
	arr := [32]byte{}
	_ = copy(arr[:], str)
	return arr
}
