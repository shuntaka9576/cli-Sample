package main

import (
	"github.com/urfave/cli"
	"os"
)

var Vesion string = "0.1.0"

func main() {
	newApp().Run(os.Args)

}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "cli-Sample"
	app.Usage = "explain cli tool"
	app.Version = Vesion
	app.Author = "shuntaka9576"
	app.Email = ""
	app.Commands = Commands
	return app
}
