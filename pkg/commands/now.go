package commands

import (
	"flag"
	"fmt"
	"os"
	"path"

	"goweather.com/goweather/internal/config"
	"goweather.com/goweather/pkg/api"
)

var nowConfig struct {
	city     string
	detailed bool
}

var nowUsage = `
Usage: %s now [options] [city name]
Get Weather data for city

Options:
`

func GetCurrentWeather(cfg *config.Config, args []string) {

	fs := flag.NewFlagSet("now", flag.ExitOnError)
	fs.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, nowUsage, name)
		fs.PrintDefaults()
	}
	fs.StringVar(&nowConfig.city, "city", "Dublin", "Name of the city")
	fs.BoolVar(&nowConfig.detailed, "detailed", false, "Detailed description of weather")
	fs.Parse(args[1:])

	geocode, gErr := api.GetGeoCode(cfg.BaseUrlGeoCoding, cfg.APIKey, nowConfig.city)

	if gErr != nil {
		fmt.Fprintf(os.Stderr, "Error fetching GeoCode: %s", gErr)
	}

	out, err := api.GetWeatherByCity(cfg.APIKey, cfg.BaseUrlCurrent, geocode, nowConfig.detailed)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
	}

	fmt.Fprintf(os.Stdout, out)
}
