package main

import (
	"os"

	"github.com/uussoop/idp-go/pkg/services"
)

func main() {
	services.InitSqlDB()
	if os.Getenv("DEBUG") != "TRUE" {
		services.InitPairKeysAndProviders()
		services.InitCron()
	}
	services.InitBlocks()
	services.InitGin()

}
