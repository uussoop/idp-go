package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uussoop/idp-go/database"
	"github.com/uussoop/idp-go/pkg/customerrors"
	"github.com/uussoop/idp-go/pkg/jwt"
	"github.com/uussoop/idp-go/pkg/utils"
)

type LoginRequest struct {
	Address string `json:"address"`
	Nonce   string `json:"nonce"`
	Sig     string `json:"sig"`
}
type LoginResponse struct {
	AccessToken string `json:"token"`
}

func (s LoginRequest) Validate() *customerrors.Customerror {
	if !hexRegex.MatchString(s.Address) {
		return &customerrors.InvalidAddress
	}
	if !nonceRegex.MatchString(s.Nonce) {
		return &customerrors.NoncCheck
	}
	if len(s.Sig) == 0 {
		return &customerrors.SigMissing
	}
	return nil
}

func LoginHandler(c *gin.Context) {

	var p LoginRequest
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.BadBodyRequest)
		return
	}
	if err := p.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, *err)
		return
	}
	address := strings.ToLower(p.Address)
	user, err := Authenticate(address, p.Nonce, p.Sig)
	if err != nil {
		c.JSON(http.StatusBadRequest, *err)
		return
	}

	signedToken, jwterr := jwt.CreateToken(user.Address, time.Minute*15)
	if jwterr != nil {
		c.JSON(http.StatusBadRequest, customerrors.CreateCustomError(500, "failed in making jwt"))
		return
	}
	// trigger sending the signed jwt to sp's
	//
	c.JSON(http.StatusOK, LoginResponse{AccessToken: signedToken})

}

func Authenticate(
	address string,
	nonce string,
	sigHex string,
) (database.User, *customerrors.Customerror) {
	user, err := database.GetUserByAddress(address)
	if err != nil {
		return *user, &customerrors.UserNotFound
	}
	if user.Nonce != nonce {
		return *user, &customerrors.NoncCheck
	}

	sig := hexutil.MustDecode(sigHex)
	// https://github.com/ethereum/go-ethereum/blob/master/internal/ethapi/api.go#L516
	// check here why I am subtracting 27 from the last byte
	logrus.Warn("sig : ", sig)
	sig[crypto.RecoveryIDOffset] -= 27
	msg := accounts.TextHash([]byte(nonce))
	logrus.Warn("msg : ", msg)

	recovered, err := crypto.SigToPub(msg, sig)
	logrus.Warn("recovered : ", recovered)
	if err != nil {
		return *user, &customerrors.SigToPub
	}
	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	logrus.Warn("recovered address", recoveredAddr.Hex())

	if user.Address != strings.ToLower(recoveredAddr.Hex()) {
		return *user, &customerrors.AuthFailure
	}

	// update the nonce here so that the signature cannot be resused
	nonce, err = utils.GetNonce()
	if err != nil {
		return *user, &customerrors.NonceGeneration
	}
	user.Nonce = nonce
	user.Update()

	return *user, nil
}
