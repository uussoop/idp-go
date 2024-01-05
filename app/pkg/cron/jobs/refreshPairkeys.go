package jobs

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/uussoop/idp-go/database"
	"github.com/uussoop/idp-go/pkg/jwt"
	"github.com/uussoop/idp-go/pkg/providers"
	"github.com/uussoop/idp-go/pkg/utils"
)

func RefreshKeys() {

	jwt.PrivateKey = utils.GeneratePairKey()
	providerslist, err := database.GetAllServiceProviders()
	if err != nil {
		logrus.Error(err)
		return
	}
	pub, err := utils.Rs512PubToByte(&(jwt.PrivateKey.PublicKey))
	if err != nil {
		logrus.Error(err)
		return
	}
	//push to service providers
	failed, err := providers.PushKeyToProviders(pub, providerslist)
	retryqouta := 20
	for i := 1; i <= retryqouta; i++ {
		if err == nil && len(failed) > 0 {
			time.Sleep(5 * time.Minute)
			failed, err = providers.PushKeyToProviders(pub, failed)

		}
	}
}
