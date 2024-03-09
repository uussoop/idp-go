package database

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func InitDB() (err error) {
	uri := os.Getenv("MYSQL_URI")

	gCnf := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	if uri == "" {
		log.Info("Initializing sqlite database...")
		// DB, err = gorm.Open(sqlite.Open("database/sqlite.db"), gCnf) /app/database/
		err = os.Mkdir("database", 0777)
		if err != nil {
			fmt.Println(err)
		}

		DB, err = openDB("./sqlite.db", gCnf)

	} else {

		DB, err = gorm.Open(mysql.Open(uri+"?parseTime=true"), gCnf)

	}
	if err != nil {
		log.Info("failed to connect database: ", err)
	}

	err = DB.AutoMigrate(
		// add models to migrate
		&User{},
		&ServiceProviders{},
		&UserWhitelist{},
	)
	log.Info("error migration ", err)

	log.Info("SQL Database has been initialized")
	return

}
func openDB(filepath string, gCnf *gorm.Config) (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(filepath), gCnf)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
