package api

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uussoop/idp-go/database"
	"github.com/uussoop/idp-go/pkg/customerrors"
	"github.com/uussoop/idp-go/pkg/utils"
)

var (
	hexRegex   *regexp.Regexp = regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	nonceRegex *regexp.Regexp = regexp.MustCompile(`^[0-9]+$`)
)

type RegisterRequest struct {
	Address  string `json:"address"`
	Username string `json:"username"`
}

func (p RegisterRequest) Validate() *customerrors.Customerror {
	if !hexRegex.MatchString(p.Address) {
		return &customerrors.InvalidAddress
	}
	return nil
}

func RegisterHandler(c *gin.Context) {

	var p RegisterRequest
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.BadBodyRequest)
		return
	}
	if err := p.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, *err)
		return
	}
	nonce, err := utils.GetNonce()
	if err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NonceGeneration)
		return
	}
	u := database.User{
		Address:  strings.ToLower(p.Address), // let's only store lower case
		Username: p.Username,
		Nonce:    nonce,
	}
	if err := u.Create(); err != nil {
		logrus.Warn(err)
		c.JSON(http.StatusInternalServerError, customerrors.Usercreation)
		return
	}
	c.JSON(http.StatusOK, customerrors.Success)

}
