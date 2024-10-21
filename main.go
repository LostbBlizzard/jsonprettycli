package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/tidwall/pretty"
)

var CLI struct {
	Fmt struct {
		Input string `help:"input File"`
	} `cmd:""  help:"formats the json file"  type:"path"`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "fmt":
		file, err := os.ReadFile(CLI.Fmt.Input)
		if err != nil {
			panic(err)
		}

		result := pretty.Pretty(file)
		os.WriteFile(CLI.Fmt.Input, result, 0644)

	default:
		panic(ctx.Command())
	}
}
