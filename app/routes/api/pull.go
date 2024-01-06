package api

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uussoop/idp-go/database"
	"github.com/uussoop/idp-go/pkg/customerrors"
	"github.com/uussoop/idp-go/pkg/jwt"
	"github.com/uussoop/idp-go/pkg/utils"
)

type PullResponse struct {
	PublickKey string `json:"public_key"`
}

func PullHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	ip := c.ClientIP()
	logrus.Info(token)
	logrus.Info(ip)

	if token == "" {
		c.JSON(
			http.StatusUnauthorized,
			customerrors.Unauthorized,
		)
		c.Abort()
		return
	} else {
		token = strings.Replace(token, "Bearer ", "", 1)
		err := database.GetByIpAndToken(ip, token)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized,
				customerrors.Unauthorized,
			)
			c.Abort()
			return
		}
		pub, err := utils.Rs512PubToByte(&(jwt.PrivateKey.PublicKey))
		pstring := base64.StdEncoding.EncodeToString(pub)
		if err != nil {
			logrus.Error(err)
			c.JSON(
				http.StatusInternalServerError,
				customerrors.InvalidAddress,
			)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, PullResponse{PublickKey: pstring})

	}

}
