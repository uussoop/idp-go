package services

import (
	"errors"
	"fmt"
	"os"

	"github.com/uussoop/idp-go/routes"
)

func InitGin() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println(errors.New("please set PORT environment variable"))
		port = "8080"
		fmt.Println("Defaulting to port ", port)

	}

	err := routes.InitRouter().Run(":" + port)
	if err != nil {
		return
	}

}
