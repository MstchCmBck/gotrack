package provider

import (
	"github.com/mstch/gotrack/utils/language"
	"github.com/mstch/gotrack/utils/json"
)

type Request interface {
	Send() ([]byte, error)
	Result()
}

type newRequestBuilderFunc func() *requestBuilder

type requestBuilder interface {
	addToken(string) requestBuilder
	addPackageId(string) requestBuilder
	addLanguage(string) requestBuilder
	build() Request
}

type requestBuilderFunc func() requestBuilder

func NewRequest(provider, packageId string) Request {
	lang, _ := language.Get()

	var request Request
	config := json.Parse()
	for _, providerFromConfig := range config.Providers {
		// Verify if provider from the command line match a provider from the config file
		if (provider != providerFromConfig.Name) {
			continue
		}

		request = builder(provider)().addToken(providerFromConfig.APIToken).
			addPackageId(packageId).
			addLanguage(lang).
			build()
	}

	return request
}

func builder(provider string) requestBuilderFunc {
	if provider == "laposte" {
		return newLaPosteRequestBuilder
	}
	return nil
}
