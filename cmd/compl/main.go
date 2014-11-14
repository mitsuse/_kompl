package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/compl"
)

func main() {
	app := initApp()
	app.Run(os.Args)
}

func initApp() *cli.App {
	app := cli.NewApp()

	app.Name = "compl"
	app.Version = "0.0.1"
	app.Usage = "A server for N-gram based word completion implemented in Golang."
	app.Action = execute

	return app
}

func execute(context *cli.Context) {
	s := compl.NewServer()

	if err := s.Run(); err != nil {
		// TODO: Handle an error.
		return
	}
}
