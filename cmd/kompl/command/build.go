package command

import (
	"compress/gzip"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/kompl/predictor"
)

func NewBuildCommand() cli.Command {
	command := cli.Command{
		Name:      "build",
		ShortName: "b",
		Usage:     "Builds a word predictor for the Kompl server.",
		Action:    buildAction,

		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "ngram-order,n",
				Value: 3,
				Usage: "The order of N-gram.",
			},

			cli.StringFlag{
				Name:  "predictor,p",
				Value: "predictor.kompl",
				Usage: "The output path of a word predictor.",
			},

			cli.StringFlag{
				Name:  "raw,r",
				Value: "predictor.kompl",
				Usage: "The input path of a word-segmented corpus.",
			},
		},
	}

	return command
}

func buildAction(context *cli.Context) {
	order := context.Int("ngram-order")
	if order < 2 {
		PrintError(ERROR_INVALID_NGRAM_ORDER, nil)
		return
	}

	rawFile, err := os.Open(context.String("raw"))
	if err != nil {
		PrintError(ERROR_LOADING_CORPUS, err)
		return
	}
	defer rawFile.Close()

	p, err := predictor.Build(order, rawFile)
	if err != nil {
		PrintError(ERROR_BUILDING_PREDICTOR, err)
		return
	}

	predictorFile, err := os.Create(context.String("predictor"))
	if err != nil {
		PrintError(ERROR_WRITING_PREDICTOR, err)
		return
	}

	defer predictorFile.Close()

	gzipWriter := gzip.NewWriter(predictorFile)
	defer gzipWriter.Close()

	if err := predictor.Dump(p, gzipWriter); err != nil {
		PrintError(ERROR_WRITING_PREDICTOR, err)
		return
	}
}
