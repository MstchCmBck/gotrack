package provider

import (
	"github.com/mstch/gotrack/utils/language"
)

type Request interface {
	Send() ([]byte, error)
	Result()
}

type requestBuilder interface {
	addToken()
	addPackageId()
	addLanguage()
	build()
}

// TODO replace this hard coded value by something store in a config file
const (
	token = "TuJRyLm1pYNrM+p+9rNLd4/ZeIpYpAD4Abma3ot2g0llimorjYNfF338D4grlAWy"
)

func NewRequest(provider, packageId string) Request {
	lang, _ := language.Get()

	var request Request
	if (provider == "laposte") {

		request = newLaPosteRequestBuilder().addToken(token).
			addPackageId(packageId).
			addLanguage(lang).
			build()
	}

	return request
}
