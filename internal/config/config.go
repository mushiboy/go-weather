package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	APIKey           string `json:"api_key"`
	BaseUrlCurrent   string `json:"base_url_current"`
	BaseUrlForecast  string `json:"base_url_forecast"`
	BaseUrlGeoCoding string `json:"base_url_geocoding"`
}

func LoadConfig() (*Config, error) {
	// filepath := "./internal/config/config.json"
	filepath := "/usr/local/etc/goweather/internal/config/config.json"
	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	config := &Config{}
	err = decoder.Decode(config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
