package services

import (
	"github.com/sirupsen/logrus"
	"github.com/uussoop/idp-go/pkg/jwt"
	"github.com/uussoop/idp-go/pkg/providers"
	"github.com/uussoop/idp-go/pkg/utils"
)

func InitPairKeysAndProviders() {
	jwt.PrivateKey = utils.GeneratePairKey()

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
