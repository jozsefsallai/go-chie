package commands

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jozsefsallai/go-chie"
	"github.com/urfave/cli"
)

func handle(action string, ctx *cli.Context) error {
	if ctx.NArg() != 1 {
		return errors.New("too few or too many arguments provided to the command")
	}

	from := ctx.Args().Get(0)
	to := ctx.String("output")

	matches, err := filepath.Glob(from)
	if err != nil {
		return err
	}

	if len(matches) == 0 {
		return errors.New("no files match the pattern that was provided")
	}

	for _, match := range matches {
		tsc := chie.NewTSCParser()

		if action == "decrypt" {
			fmt.Printf("Decrypting \"%s\"...\n", match)
		} else {
			fmt.Printf("Encrypting \"%s\"...\n", match)
		}

		basename := filepath.Base(match)
		output := strings.Replace(to, "*", strings.TrimSuffix(basename, filepath.Ext(basename)), 1)

		err = tsc.FromFile(match)
		if err != nil {
			return err
		}

		if action == "decrypt" {
			tsc.Decrypt()
		} else if action == "encrypt" {
			tsc.Encrypt()
		}

		err = tsc.ToFile(output)
		if err != nil {
			return err
		}

		fmt.Printf("Done: %s\n", output)
	}

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
				return handle("encrypt", ctx)
			},
		},
	},
}
