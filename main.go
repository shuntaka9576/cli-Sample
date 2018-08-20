package main

import (
	"os"
	"github.com/urfave/cli"
)

var Vesion string = "0.1.0"

func main() {
	newApp().Run(os.Args)

}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "ghq"
	app.Usage = "explain cli tool"
	app.Version = Vesion
	app.Author = "shuntaka9576"
	app.Email = ""
	//app.Commands = Commands
	return app
}
