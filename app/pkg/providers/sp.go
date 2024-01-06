package providers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/uussoop/idp-go/database"
)

type pubkeyBody struct {
	PublickKey string `json:"public_key"`
}

func PushKeyToProviders(
	pubKey []byte,
	providers []database.ServiceProviders,
) (failedProviders []database.ServiceProviders, err error) {
	var pubkeyBody pubkeyBody
	pstring := base64.StdEncoding.EncodeToString(pubKey)
	pubkeyBody.PublickKey = pstring
	pubKey, err = json.Marshal(pubkeyBody)
	if err != nil {
		return nil, err
	}
	// Create POST request
	req, err := http.NewRequest("POST", "url", bytes.NewBuffer(pubKey))
	if err != nil {
		return nil, err
	}

	for _, provider := range providers {
		// Add bearer token header
		req.Header.Add("Authorization", "Bearer "+provider.Token)
		req.Header.Add("Content-Type", "application/json")
		urlParsed, err := url.Parse(provider.URL)
		if err != nil {
			provider.Description = "Invalid URL"
			failedProviders = append(failedProviders, provider)
		}
		req.URL.Host = urlParsed.Host
		req.URL.Scheme = urlParsed.Scheme
		req.URL.Path = urlParsed.Path

		// Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			provider.Description = "failed to send request"
			failedProviders = append(failedProviders, provider)
		}
		defer resp.Body.Close()

		// Check response
		if resp.StatusCode != http.StatusOK {
			provider.Description = "status code is " + resp.Status
			failedProviders = append(failedProviders, provider)
		}
	}

	return failedProviders, nil
}
