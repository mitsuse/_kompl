package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/kompl/commands"
)

func main() {
	app := initApp()
	app.Run(os.Args)
}

func initApp() *cli.App {
	app := cli.NewApp()

	app.Name = commands.NAME
	app.Version = commands.VERSION
	app.Usage = commands.DESCRIPTION

	app.Commands = []cli.Command{
		commands.NewRunCommand(),
		commands.NewBuildCommand(),
	}

	return app
}
