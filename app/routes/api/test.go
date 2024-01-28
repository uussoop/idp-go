package api

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/uussoop/idp-go/pkg/blocks"
)

type TestRequest struct {
	Address string `json:"address"`
}
type TestResponse struct {
	Balance string `json:"balance"`
}

func TestHandler(c *gin.Context) {
	var request TestRequest
	var response TestResponse
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	balance, err := blocks.BscContract.GetTokenBalance(common.HexToAddress(request.Address))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	stringbalance := fmt.Sprintf("%f", balance)
	response.Balance = stringbalance
	c.JSON(http.StatusOK, response)

}
