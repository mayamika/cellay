package main

import (
	"fmt"
	"os"

	"github.com/mayamika/cellay/cellay-server/app"
)

func main() {
	config, err := app.ParseFlagsAndConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	app.New(config).Run()
}
