package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uussoop/idp-go/pkg/jwt"
)

type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

func GetBalanceHandler(c *gin.Context) {
	tokenfull := c.GetHeader("Authorization")
	token := strings.Replace(tokenfull, "Bearer ", "", 1)

	username, ok, err := jwt.ValidateToken(token)

	response := BalanceResponse{}
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	response.Address = username
	response.Balance = "1"
	if ok {
		c.JSON(http.StatusOK, response)

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
}
