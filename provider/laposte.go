package provider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
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

	builder.request.url.Path += url.QueryEscape(builder.packageId)
	builder.request.url.RawQuery += "lang=" + url.QueryEscape(builder.lang)

	return builder.request
}

func (r LaPosteRequest) String() (string) {
	return r.url.String()
}

func (r *LaPosteRequest) send() ([]byte, error) {
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

type Event struct {
	Date  string `json:"date"`
	Label string `json:"label"`
}

type TimelineElement struct {
	ShortLabel string `json:"shortLabel"`
}

type Shipment struct {
	Events   []Event           `json:"event"`
	Timeline []TimelineElement `json:"timeline"`
}

func parseData(input []byte) []string {
	// Define the struct to store the JSON data
	type Data struct {
		Shipment Shipment `json:"shipment"`
	}

	// Parse the JSON data
	var data Data
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		logger.Error(err)
	}

	// Retrieve the event dates and labels
	var output []string
	output = append(output, "Events:")
	fmt.Println("Events:")
	for _, event := range data.Shipment.Events {
		result := fmt.Sprintf("Date: %s\nLabel: %s\n\n", event.Date, event.Label)
		output = append(output, result)
	}

	// Retrieve the last element's short label from the timeline
	lastIndex := len(data.Shipment.Timeline) - 1
	if lastIndex >= 0 {
		lastElement := data.Shipment.Timeline[lastIndex]
		result := fmt.Sprintf("Last Element Short Label:\n\t%s", lastElement.ShortLabel)
		output = append(output, result)
	}

	return output
}

func (r LaPosteRequest) GetResult() []string {
	data, _ := r.send()
	return parseData(data)
}
