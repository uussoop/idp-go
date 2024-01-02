package services

import (
	log "github.com/sirupsen/logrus"
	"github.com/uussoop/idp-go/database"
)

func InitSqlDB() {
	err := database.InitDB()
	if err != nil {
		log.Error("Error: ", err)
		panic(err)
	}

}
