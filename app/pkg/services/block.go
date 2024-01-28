package services

import (
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uussoop/idp-go/pkg/blocks"
)

func InitBlocks() {
	clientUrl := os.Getenv("BLOCKS_CLIENT_URL")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	contractName := os.Getenv("CONTRACT_NAME")
	contractSymbol := os.Getenv("CONTRACT_SYMBOL")
	contractDecimal := os.Getenv("CONTRACT_DECIMAL")
	contractTotalSupply := os.Getenv("CONTRACT_TOTALSUPPLY")
	balanceLimit := os.Getenv("BALANCE_LIMIT")

	if clientUrl == "" || contractAddress == "" || contractName == "" || contractSymbol == "" ||
		contractDecimal == "" ||
		contractTotalSupply == "" {
		panic("couldnt init the block client")
	}
	var err error
	blocks.BalanceLimit, err = strconv.ParseFloat(balanceLimit, 64)
	blocks.Client, err = ethclient.Dial(clientUrl)
	blocks.Url = clientUrl

	contractdecimalint, err := strconv.ParseInt(contractDecimal, 10, 64)
	contractTotalSupplyint, err := strconv.ParseInt(contractTotalSupply, 10, 64)
	if err != nil {
		panic("couldnt init the block client")
	}

	blocks.BscContract = &blocks.Contract{
		Ethereum: &blocks.Ethereum{
			Client: blocks.Client,
		},
		Address:     common.HexToAddress(contractAddress),
		Name:        contractName,
		Symbol:      contractSymbol,
		Decimals:    int(contractdecimalint),
		TotalSupply: int(contractTotalSupplyint),
	}

}
