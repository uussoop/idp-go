package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uussoop/idp-go/database"
	"github.com/uussoop/idp-go/pkg/customerrors"
)

type NonceRequest struct {
	Address string `json:"address"`
}
type NonceResponse struct {
	Nonce string `json:"nonce"`
}

func UserNonceHandler(c *gin.Context) {

	var r NonceRequest
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.BadBodyRequest)
		return
	}
	if !hexRegex.MatchString(r.Address) {
		c.JSON(http.StatusBadRequest, customerrors.NoncCheck)
		return
	}
	user, err := database.GetUserByAddress(strings.ToLower(r.Address))
	if err != nil {
		c.JSON(http.StatusForbidden, customerrors.UserNotFound)

		return
	}
	c.JSON(http.StatusOK, NonceResponse{Nonce: user.Nonce})
}
