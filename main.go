package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/web3art/g/cmd"
)

func main() {
	app := &cli.App{
		Name:  "serve",
		Usage: "Start dapr miro service",
		Action: func(c *cli.Context) error {
			return cmd.Serve(c)
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
