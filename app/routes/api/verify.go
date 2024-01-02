package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uussoop/idp-go/pkg/customerrors"
	"github.com/uussoop/idp-go/pkg/jwt"
)

type VerifyRequest struct {
	AccessToken string `json:"token"`
}

func VerifyHandler(c *gin.Context) {
	var p VerifyRequest
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.BadBodyRequest)
		return
	}
	if p.AccessToken == "" {
		c.JSON(http.StatusBadRequest, customerrors.BadBodyRequest)
		return
	}
	tokenString := p.AccessToken
	if len(tokenString) == 0 {
		c.JSON(http.StatusUnauthorized, customerrors.Unauthorized)
		return
	}

	ok, err := jwt.ValidateToken(tokenString)
	if err != nil || !ok {
		c.JSON(http.StatusUnauthorized, customerrors.Unauthorized)
		return
	}
	c.JSON(http.StatusOK, customerrors.SuccessVerify)

}
