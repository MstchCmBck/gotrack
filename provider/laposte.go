package provider

import (
	"net/url"
	"io/ioutil"
	"net/http"
	"github.com/mstch/gotrack/utils/logger"
)

type LaPosteRequest struct {
	token string
	url url.URL
}

type laPosteRequestBuilder struct {
	packageId string
	lang string
	request *LaPosteRequest
}

func newLaPosteRequestBuilder() requestBuilder {
	return &laPosteRequestBuilder{request: &LaPosteRequest{}}
}

func (builder *laPosteRequestBuilder) addToken(t string) requestBuilder {
	builder.request.token = t
	return builder
}

func (builder *laPosteRequestBuilder) addPackageId(packageId string) requestBuilder {
	builder.packageId = packageId
	return builder
}

func (builder *laPosteRequestBuilder) addLanguage(lang string) requestBuilder {
	builder.lang = lang
	return builder
}

func (builder *laPosteRequestBuilder) build() Request {
	builder.request.url.Scheme = "https"
	builder.request.url.Host = "api.laposte.fr"
	builder.request.url.Path = "/suivi/v2/idships/"

	builder.request.url.Path = url.QueryEscape(builder.packageId)
	builder.request.url.RawQuery += url.QueryEscape(builder.lang)

	return builder.request
}

func (r LaPosteRequest) String() (string) {
	return r.url.String()
}

func (r *LaPosteRequest) Send() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, r.url.String(), nil)
	if err != nil {
		logger.Error("creation of GET request failed.")
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Okapi-Key", r.token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("request failed")
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("ioutil.Readall fail")
		return nil, err
	}

	return data, nil
}

func (r LaPosteRequest) Result() {

}
