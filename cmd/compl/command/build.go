package command

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/compl"
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
				Usage: "The input path of a raw count file.",
			},
		},
	}

	return command
}

func buildAction(context *cli.Context) {
	// TODO: Convert a raw count file into binary formatted model.
	arpaFile, err := os.Open(context.String("arpa"))
	if err != nil {
		// TODO: Handle an error.
		return
	}
	defer arpaFile.Close()

	model, err := compl.InflateRawModel(arpaFile)
	if err != nil {
		// TODO: Handle an error.
		return
	}

	modelFile, err := os.Create(context.String("model"))
	if err != nil {
		// TODO: Handle an error.
		return
	}
	defer modelFile.Close()

	if err := model.Deflate(modelFile); err != nil {
		// TODO: Handle an error.
		return
	}
}
