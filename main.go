/*
Command "kompl" is a tool to build or run a server
for K-best word completion based on N-gram frequency.
*/
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

	app.Author = commands.AUTHOR
	app.Email = commands.AUTHOR_EMAIL

	app.Commands = []cli.Command{
		commands.NewRunCommand(),
		commands.NewBuildCommand(),
	}

	return app
}
