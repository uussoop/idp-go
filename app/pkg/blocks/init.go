package blocks

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

var BscContract *Contract
var Client *ethclient.Client
var Url string

var BalanceLimit float64

type Ethereum struct {
	Client *ethclient.Client
}

type Contract struct {
	*Ethereum
	Address     common.Address
	Name        string
	Symbol      string
	Decimals    int
	TotalSupply int
}

func (c *Contract) GetTokenBalance(accountAddress common.Address) (float64, error) {
	data := MethodPack("balanceOf(address)", common.LeftPadBytes(accountAddress.Bytes(), 32))
	msg := ethereum.CallMsg{From: common.Address{}, To: &c.Address, Data: data}
	result, err := c.Client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return 0, err
	}
	decimals, err := c.GetTokenDecimals()
	if err != nil {
		return 0, err
	}
	return BigIntToAmount(common.BytesToHash(result[:]).Big(), decimals), nil
}

func BigIntToAmount(val *big.Int, decimals int) float64 {
	bigval := new(big.Float)
	bigval.SetInt(val)
	coin := new(big.Float)
	dec := new(big.Int)
	dec.SetInt64(int64(decimals))
	coin.SetInt(big.NewInt(10).Exp(big.NewInt(10), dec, nil))
	bigval.Quo(bigval, coin)
	result, _ := bigval.Float64()
	return result
}
func MethodPack(method string, args ...[]byte) []byte {
	fnSignature := []byte(method)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(fnSignature)
	methodID := hash.Sum(nil)[:4]
	var data []byte
	data = append(data, methodID...)
	for _, arg := range args {
		data = append(data, arg...)
	}
	return data
}
func (c *Contract) GetTokenDecimals() (int, error) {

	if c.Decimals != 0 {
		return c.Decimals, nil
	}
	data := MethodPack("decimals()")
	msg := ethereum.CallMsg{From: common.Address{}, To: &c.Address, Data: data}
	result, err := c.Client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return 0, err
	}
	return int(common.BytesToHash(result[:]).Big().Int64()), nil
}
