package json

import (
	"encoding/json"
	"io/ioutil"
	"github.com/mstch/gotrack/utils/logger"
)

type Provider struct {
	Name     string `json:"name"`
	APIToken string `json:"apiToken"`
}

type ProviderConfig struct {
	Providers []Provider `json:"providers"`
}

func Parse() ProviderConfig {
	// Read JSON file
	data, err := ioutil.ReadFile("providers.json")
	if err != nil {
		logger.Error("Error reading JSON file:", err)

	}

	// Parse JSON data
	var config ProviderConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		logger.Error("Error parsing JSON:", err)
	}

	return config
}
