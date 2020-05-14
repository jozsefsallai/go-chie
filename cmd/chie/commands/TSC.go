package commands

import (
	"errors"
	"fmt"

	"github.com/jozsefsallai/go-chie"
	"github.com/urfave/cli"
)

func handle(action string, ctx *cli.Context) error {
	if ctx.NArg() != 1 {
		return errors.New("too few or too many arguments provided to the command")
	}

	from := ctx.Args().Get(0)
	to := ctx.String("output")

	tsc := chie.NewTSCParser()

	err := tsc.FromFile(from)
	if err != nil {
		return err
	}

	if action == "decrypt" {
		tsc.Decrypt()
	} else if action == "encrypt" {
		tsc.Encrypt()
	}

	err = tsc.ToFile(to)
	if err != nil {
		return err
	}

	fmt.Println("Done!")
	return nil
}

// TSCCommand is the command that encrypts/decrypts TSC
var TSCCommand = cli.Command{
	Name:  "tsc",
	Usage: "Convert to/from TSC",
	Subcommands: []cli.Command{
		{
			Name:  "decrypt",
			Usage: "decrypt a TSC file into human-readable text",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Value: "output.txt",
					Usage: "`PATH` to output file",
				},
			},
			Action: func(ctx *cli.Context) error {
				fmt.Println("Decrypting...")
				return handle("decrypt", ctx)
			},
		},
		{
			Name:  "encrypt",
			Usage: "encrypt a plaintext file into TSC",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Value: "output.tsc",
					Usage: "`PATH` to output file",
				},
			},
			Action: func(ctx *cli.Context) error {
				fmt.Println("Encrypting...")
				return handle("encrypt", ctx)
			},
		},
	},
}
