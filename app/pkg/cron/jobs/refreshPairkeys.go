package jobs

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/uussoop/idp-go/database"
	"github.com/uussoop/idp-go/pkg/jwt"
	"github.com/uussoop/idp-go/pkg/providers"
	"github.com/uussoop/idp-go/pkg/utils"
)

func RefreshKeys() {
	privkeypath := os.Getenv("PRIVATE_KEY_PATH")

	jwt.PrivateKey = utils.GeneratePairKey()
	prbyte, err := utils.Rs512PrivToByte(jwt.PrivateKey)
	err = os.WriteFile(privkeypath, prbyte, 0644)
	if err != nil {
		logrus.Error(err)
	}
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
