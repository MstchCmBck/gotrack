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

type LaPosteRequestBuilder struct {
	packageId string
	lang string
	request *LaPosteRequest
}

func NewLaPosteRequestBuilder() *LaPosteRequestBuilder {
	return &LaPosteRequestBuilder{request: &LaPosteRequest{}}
}

func (builder *LaPosteRequestBuilder) AddToken(t string) *LaPosteRequestBuilder {
	builder.request.token = t
	return builder
}

func (builder *LaPosteRequestBuilder) AddPackageId(packageId string) *LaPosteRequestBuilder {
	builder.packageId = packageId
	return builder
}

func (builder *LaPosteRequestBuilder) AddLanguage(lang string) *LaPosteRequestBuilder {
	builder.lang = lang
	return builder
}

func (builder *LaPosteRequestBuilder) Build() *LaPosteRequest {
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
