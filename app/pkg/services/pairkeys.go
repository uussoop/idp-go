package services

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/uussoop/idp-go/pkg/jwt"
	"github.com/uussoop/idp-go/pkg/providers"
	"github.com/uussoop/idp-go/pkg/utils"
)

func InitPairKeysAndProviders() {
	privkeypath := os.Getenv("PRIVATE_KEY_PATH")
	if priv, err := os.ReadFile(privkeypath); len(priv) != 0 && err == nil {

		var err error

		jwt.PrivateKey, err = utils.ByteToRs512Priv([]byte(priv))
		if err != nil {
			logrus.Warn("invalid priv key  ", err)
			jwt.PrivateKey = utils.GeneratePairKey()
			prbyte, err := utils.Rs512PrivToByte(jwt.PrivateKey)
			err = os.WriteFile(privkeypath, prbyte, 0644)
			if err != nil {
				panic("error in creating and saving private keys ")
			}
		}
	} else {
		jwt.PrivateKey = utils.GeneratePairKey()
		prbyte, err := utils.Rs512PrivToByte(jwt.PrivateKey)
		err = os.WriteFile(privkeypath, prbyte, 0644)
		if err != nil {
			panic("error in creating and saving private keys ")
		}
	}

	c, err := providers.ReadProviderConfig("./config/config.yaml")
	if err != nil {
		logrus.Error(err)
		return
	}
	for _, u := range c.Whitelist {
		err = u.Create()
		if err != nil {
			logrus.Error(err)

		}
	}

	for _, v := range c.Providers {
		err = v.Create()
		if err != nil {
			logrus.Error(err)

		}
	}
	pub, err := utils.Rs512PubToByte(&(jwt.PrivateKey.PublicKey))
	if err != nil {
		logrus.Error(err)
		return
	}
	providers.PushKeyToProviders(pub, c.Providers)

}
