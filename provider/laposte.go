package provider

import (
	"net/url"
	"io/ioutil"
	"net/http"
)

type LaPosteRequest struct {
	token string
	url url.URL
}

func (r * LaPosteRequest) Build() {
	r.url.Scheme = "https"
	r.url.Host = "api.laposte.fr"
	r.url.Path = "/suivi/v2/idships/"
}

func (r *LaPosteRequest) AddToken(t string) {
	r.token = t
}

func (r *LaPosteRequest) AddPackageId(packageId string) {
	r.url.Path += url.QueryEscape(packageId)
}

func (r *LaPosteRequest) AddLanguage(language string) {
	r.url.RawQuery += url.QueryEscape(language)
}

func (r LaPosteRequest) String() (string) {
	return r.url.String()
}

func (r *LaPosteRequest) Send() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, r.url.String(), nil)
	if err != nil {
		// logger.Error("creation of GET request failed.")
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Okapi-Key", r.token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// logger.Error("request failed")
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// logger.Error("ioutil.Readall fail")
		return nil, err
	}

	return data, nil
}
