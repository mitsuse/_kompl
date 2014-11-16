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
	app.Usage = "A server for N-gram based word completion."

	app.Action = execute

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "model,m",
			Value: "model.compl",
			Usage: "The path of an N-gram completion model.",
		},

		cli.StringFlag{
			Name:  "port,p",
			Value: "8080",
			Usage: "The port number which a compl server uses.",
		},
	}

	return app
}

func execute(context *cli.Context) {
	model, err := compl.InflateModel(context.String("model"))
	if err != nil {
		// TODO: Handle an error.
		return
	}

	s := compl.NewServer(context.String("port"), model)
	if err := s.Run(); err != nil {
		// TODO: Handle an error.
		return
	}
}
