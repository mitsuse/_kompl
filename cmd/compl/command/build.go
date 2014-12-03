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
		Usage:     "Builds a binary-formatted word predictor",
		Action:    buildAction,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "predictor,p",
				Value: "predictor.compl",
				Usage: "The output path of a word predictor.",
			},

			cli.StringFlag{
				Name:  "raw,r",
				Value: "predictor.compl",
				Usage: "The input path of a raw count file.",
			},
		},
	}

	return command
}

func buildAction(context *cli.Context) {
	// TODO: Convert a raw count file into binary formatted predictor.
	rawFile, err := os.Open(context.String("raw"))
	if err != nil {
		// TODO: Handle an error.
		return
	}
	defer rawFile.Close()

	predictor, err := compl.InflateRawPredictor(rawFile)
	if err != nil {
		// TODO: Handle an error.
		return
	}

	predictorFile, err := os.Create(context.String("predictor"))
	if err != nil {
		// TODO: Handle an error.
		return
	}
	defer predictorFile.Close()

	if err := predictor.Deflate(predictorFile); err != nil {
		// TODO: Handle an error.
		return
	}
}
