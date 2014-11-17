package command

import (
	"github.com/codegangsta/cli"
	"github.com/mitsuse/compl"
)

func NewRunCommand() cli.Command {
	command := cli.Command{
		Name:      "run",
		ShortName: "r",
		Usage:     "Runs a seal server.",
		Action:    runAction,

		Flags: []cli.Flag{
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
		},
	}

	return command
}

func runAction(context *cli.Context) {
	// TODO: Start a seal server.
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
