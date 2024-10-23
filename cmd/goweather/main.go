package main

import (
	"fmt"
	"os"
	"path"

	"goweather.com/goweather/internal/config"
	"goweather.com/goweather/pkg/commands"
)

type Command struct {
	name    string
	desc    string
	execute func(*config.Config, []string)
}

func (c Command) String() string {
	return fmt.Sprintf("%10s : %s", c.name, c.desc)
}

var cs = map[string]Command{
	"now":      {"now", "Get current weather", commands.GetCurrentWeather},
	"forecast": {"forecast", "Get forecast", commands.GetForecast},
}

func stringify(cs map[string]Command) string {
	out := ""
	for c, _ := range cs {
		out += fmt.Sprintf("%s\n", c)
	}
	return out
}

var usage = `
Usage: %s <command> [<args>]
Run %s <command> -h for sub command help
Commands:
%s
`

func usageMessage() {
	name := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, usage, name, name, stringify(cs))
}
func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config file: %s", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: missing command\n")
		usageMessage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-h", "--help":
		usageMessage()
		os.Exit(1)
	case "now", "forecast":
		command := cs[os.Args[1]]
		command.execute(cfg, os.Args[1:])
	default:
		fmt.Fprintf(os.Stderr, "Error: Wrong Command")
		usageMessage()
		os.Exit(1)
	}
}
