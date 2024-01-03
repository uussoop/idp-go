package jobs

import (
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
	providers.PushKeyToProviders(pub, providerslist)
}
