package main

import "github.com/uussoop/idp-go/pkg/services"

func main() {
	services.InitSqlDB()
	services.InitPairKeysAndProviders()
	services.InitCron()
	services.InitGin()

}
