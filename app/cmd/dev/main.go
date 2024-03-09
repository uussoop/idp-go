package main

import (
	"os"

	"github.com/uussoop/idp-go/pkg/services"
)

func main() {
	os.Setenv("BLOCKS_CLIENT_URL", "https://bsc-dataseed.binance.org/")
	os.Setenv("CONTRACT_ADDRESS", "")
	os.Setenv("CONTRACT_NAME", "")
	os.Setenv("CONTRACT_SYMBOL", "")
	os.Setenv("CONTRACT_DECIMAL", "")
	os.Setenv("CONTRACT_TOTALSUPPLY", "")
	os.Setenv("BALANCE_LIMIT", "")
	os.Setenv("DEBUG", "FALSE")
	os.Setenv("PRIVATE_KEY_PATH", "")
	services.InitSqlDB()

	services.InitPairKeysAndProviders()

	services.InitBlocks()
	services.InitGin()

}
