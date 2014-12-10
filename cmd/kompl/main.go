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

	app.Name = "kompl"
	app.Version = "0.0.1"
	app.Usage = "A server for K-best word completion based on N-gram frequency."

	app.Commands = []cli.Command{
		command.NewRunCommand(),
		command.NewBuildCommand(),
	}

	return app
}
