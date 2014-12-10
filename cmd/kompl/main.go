package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/kompl/cmd/kompl/command"
)

func main() {
	app := initApp()
	app.Run(os.Args)
}

func initApp() *cli.App {
	app := cli.NewApp()

	app.Name = command.Name
	app.Version = command.Version
	app.Usage = command.Description

	app.Commands = []cli.Command{
		command.NewRunCommand(),
		command.NewBuildCommand(),
	}

	return app
}
