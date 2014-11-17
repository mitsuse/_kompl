package command

import (
	"github.com/codegangsta/cli"
)

func NewBuildCommand() cli.Command {
	command := cli.Command{
		Name:      "build",
		ShortName: "b",
		Usage:     "Builds a binary-formatted model for completion",
		Action:    buildAction,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "model,m",
				Value: "model.compl",
				Usage: "The output path of an N-gram completion model.",
			},

			cli.StringFlag{
				Name:  "arpa,a",
				Value: "model.arpa",
				Usage: "The input path of an ARPA-formatted N-gram model.",
			},
		},
	}

	return command
}

func buildAction(context *cli.Context) {
	// TODO: Convert an ARPA file into binary formatted model.
}
