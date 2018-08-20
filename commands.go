package main

import (
	"github.com/urfave/cli"
	"fmt"
)

var suffix string

var Commands = []cli.Command{
	{
		Name:  "hello",
		Usage: "if use set -t or --text",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "text, t",
				Value: "world",
				Usage: "hello xxx text",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello %s %s\n", c.String("text"), suffix)
			return nil
		},
	},
}
