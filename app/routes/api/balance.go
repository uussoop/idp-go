package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/uussoop/idp-go/pkg/blocks"
	"github.com/uussoop/idp-go/pkg/cache"
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
	getCache := cache.GetCache()
	if balance, ok := getCache.Get(fmt.Sprintf("balance_get%s", username)); ok {

		response.Balance = fmt.Sprintf("%f", balance)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Address = username
	balance, err := blocks.BscContract.GetTokenBalance(common.HexToAddress(username))
	if err != nil {
		response.Balance = "0"

	} else {

		response.Balance = fmt.Sprintf("%f", balance)
		getCache.Set(fmt.Sprintf("balance_get%s", username), fmt.Sprintf("%f", balance), 5*time.Minute)
	}
	if ok {
		c.JSON(http.StatusOK, response)

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
}
