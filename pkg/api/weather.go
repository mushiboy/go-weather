package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"goweather.com/goweather/pkg/models"
)

// func getGeoCode(apiKey string, baseUrl string) (string, error) {

// }

func GetWeatherByCity(apiKey string, baseUrl string, geoCode models.Geocode, detailed bool) (string, error) {

	url := fmt.Sprintf("%slat=%f&lon=%f&units=metric&appid=%s", baseUrl, geoCode.Lat, geoCode.Lon, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var weatherResponse models.CurrentWeather

	unwrapErr := json.Unmarshal(data, &weatherResponse)

	if unwrapErr != nil {
		return "", unwrapErr
	}

	out := weatherOut(weatherResponse, detailed)

	return out, nil
}

var detailedWeather = `
Temperature: %.2f°C
Feels Like: %.2f°C
Min Temperature: %.2f°C
Max Temperature: %.2f°C
Pressure: %.2f hPa
Humidity: %.2f percent
Wind Speed: %.2f m/s
Wind Direction: %d°
Cloudiness: %d percent
Weather: %s - %s
`

var weather = `
Temperature: %.2f°C
Feels Like: %.2f°C
Weather: %s - %s
`

func weatherOut(weatherResponse models.CurrentWeather, detailed bool) string {

	if !detailed {
		return fmt.Sprintf(weather,
			weatherResponse.Basic.Temp,
			weatherResponse.Basic.FeelsLike,
			weatherResponse.Weather[0].Main,
			weatherResponse.Weather[0].Desc)
	}

	return fmt.Sprintf(
		detailedWeather,
		weatherResponse.Basic.Temp,
		weatherResponse.Basic.FeelsLike,
		weatherResponse.Basic.TempMin,
		weatherResponse.Basic.TempMax,
		weatherResponse.Basic.Pressure,
		weatherResponse.Basic.Humidity,
		weatherResponse.Wind.Speed,
		weatherResponse.Wind.Deg,
		weatherResponse.Cloud.All,
		weatherResponse.Weather[0].Main,
		weatherResponse.Weather[0].Desc)
}
