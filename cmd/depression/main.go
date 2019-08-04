package main

import (
	"fmt"
	"io"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/urfave/cli"
	"go.stevenxie.me/bincode/encode"
	"go.stevenxie.me/bincode/internal/info"
)

func main() {
	app := cli.NewApp()
	app.Name = "depression"
	app.Usage = "Encode and decode using depression representation."
	app.UsageText = "depression [options] < (input) > (output)\n" +
		"   depression [options] (file)"

	app.Version = info.Version
	app.Author = "Steven Xie <hello@stevenxie.me>"
	app.HideHelp = true
	app.EnableBashCompletion = true
	app.Action = run

	// Configure flags.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "decode,d",
			Usage: "decode from depression representation",
		},
		cli.StringFlag{
			Name:  "key,k",
			Usage: "key to use for encoding",
		},
		cli.BoolFlag{
			Name:   "help,h",
			Hidden: true,
		},
	}

	// Run and exit with appropriate code.
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	if c.Bool("help") {
		cli.ShowAppHelp(c) // show help and exit
		return nil
	}
	args := c.Args()
	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "Error: expected 0 or 1 arguments, received %d.\n\n",
			len(args))
		cli.ShowAppHelpAndExit(c, 1)
	}

	// Either use stdin or read a file.
	var src io.Reader
	if len(args) == 0 {
		src = os.Stdin
	} else {
		file, err := os.Open(args[0])
		if err != nil {
			return errors.Wrap(err, "opening input file")
		}
		defer file.Close()
		src = file
	}

	// Either encode or decode.
	if c.Bool("decode") {
		dec := encode.NewDecoder(os.Stdin)
		if _, err := io.Copy(os.Stdout, dec); err != nil {
			return err
		}
	} else {
		// Determine key.
		key := c.String("key")
		if key == "" {
			key = "depression"
		}

		// Encode using key.
		enc, err := encode.NewEncoder(key, os.Stdout)
		if err != nil {
			return errors.Wrapf(err, "encoding using the key '%s'", key)
		}
		_, err = io.Copy(enc, src)
		fmt.Println() // append newline for formatting
		return err
	}

	return nil
}
