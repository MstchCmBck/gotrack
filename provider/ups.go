package provider

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/mstch/gotrack/utils/logger"
)

type upsRequest struct {
	token string
	url url.URL
}

type upsRequestBuilder struct {
	packageId string
	lang string
	request *upsRequest
}

func newUpsRequestBuilder() requestBuilder {
	return &upsRequestBuilder{request: &upsRequest{}}
}

func (builder *upsRequestBuilder) addToken(t string) requestBuilder {
	builder.request.token = t
	return builder
}

func (builder *upsRequestBuilder) addPackageId(packageId string) requestBuilder {
	builder.packageId = packageId
	return builder
}

func (builder *upsRequestBuilder) addLanguage(lang string) requestBuilder {
	builder.lang = lang
	return builder
}

func (builder *upsRequestBuilder) build() Request {
	builder.request.url.Scheme = "https"
	builder.request.url.Host = "wwwcie.ups.com"
	builder.request.url.Path = "/api/track/v1/details/"

	builder.request.url.Path += url.QueryEscape(builder.packageId)
	builder.request.url.RawQuery += "locale=" + url.QueryEscape(builder.lang)

	return builder.request
}

func (r upsRequest) String() (string) {
	return r.url.String()
}

func (r *upsRequest) send() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, r.url.String(), nil)
	if err != nil {
		logger.Error("creation of GET request failed.")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization: Bearer", r.token)
	req.Header.Set("transId", "string")
	req.Header.Set("transactionSrc", "testing")

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

func (r upsRequest) GetResult() []string {
	r.send()
	return nil
}
