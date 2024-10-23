package commands

import (
	"flag"
	"fmt"
	"os"
	"path"

	"goweather.com/goweather/internal/config"
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

	if nowConfig.detailed {
		fmt.Fprintf(os.Stdout, "Detailed Weather of %s :)\n", nowConfig.city)
	} else {
		fmt.Fprintf(os.Stdout, "Not so detailed of %s :(\n", nowConfig.city)
	}
}
