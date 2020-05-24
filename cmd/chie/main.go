package main

import (
	"log"
	"os"

	"github.com/jozsefsallai/go-chie/cmd/chie/commands"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "chie",
		Usage: "powerful TSC encryption/decryption tool",
		Commands: []cli.Command{
			commands.TSCCommand,
		},
		Version: "1.0.1",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
