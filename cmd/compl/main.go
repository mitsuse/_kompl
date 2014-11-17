package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/compl/cmd/compl/command"
)

func main() {
	app := initApp()
	app.Run(os.Args)
}

func initApp() *cli.App {
	app := cli.NewApp()

	app.Name = "compl"
	app.Version = "0.0.1"
	app.Usage = "A server for N-gram based word completion."

	app.Commands = []cli.Command{
		command.NewRunCommand(),
	}

	return app
}
