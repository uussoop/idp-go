package providers

import (
	"os"

	"github.com/uussoop/idp-go/database"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Whitelist []database.UserWhitelist    `yaml:"whitelist"`
	Providers []database.ServiceProviders `yaml:"providers"`
}

func ReadProviderConfig(filename string) (*Config, error) {
	var config Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
