package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/kompl/cmd/kompl/commands"
)

func main() {
	app := initApp()
	app.Run(os.Args)
}

func initApp() *cli.App {
	app := cli.NewApp()

	app.Name = commands.Name
	app.Version = commands.Version
	app.Usage = commands.Description

	app.Commands = []cli.Command{
		commands.NewRunCommand(),
		commands.NewBuildCommand(),
	}

	return app
}
