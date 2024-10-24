package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"goweather.com/goweather/pkg/models"
)

func GetGeoCode(baseUrl string, apiKey string, city string) (models.Geocode, error) {
	url := fmt.Sprintf("%sq=%s&appid=%s&limit=2", baseUrl, city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return models.Geocode{}, err
	}

	defer resp.Body.Close()

	data, readErr := io.ReadAll(resp.Body)

	if readErr != nil {
		return models.Geocode{}, readErr
	}

	var geoCodes []models.Geocode

	gUnwrapErr := json.Unmarshal(data, &geoCodes)

	if gUnwrapErr != nil {
		return models.Geocode{}, gUnwrapErr
	}

	geoCode := geoCodes[0]

	fmt.Fprintf(os.Stdout, "Weather details of %s in %s", geoCode.Name, geoCode.Country)

	return geoCode, nil
}
