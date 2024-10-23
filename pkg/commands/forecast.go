package commands

import (
	"fmt"
	"os"

	"goweather.com/goweather/internal/config"
)

func GetForecast(cfg *config.Config, args []string) {
	fmt.Fprintf(os.Stdout, "heyyyy\n")

}
