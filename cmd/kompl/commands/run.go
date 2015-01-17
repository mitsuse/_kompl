package commands

import (
	"compress/gzip"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/kompl"
	"github.com/mitsuse/kompl/predictor"
)

func NewRunCommand() cli.Command {
	command := cli.Command{
		Name:      "run",
		ShortName: "r",
		Usage:     "Runs the Kompl server.",
		Action:    runAction,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "predictor,p",
				Value: "predictor.kompl",
				Usage: "The path of a word predictor.",
			},

			cli.StringFlag{
				Name:  "port,n",
				Value: "8901",
				Usage: "The port number which the Kompl server uses.",
			},
		},
	}

	return command
}

func runAction(context *cli.Context) {
	predictorFile, err := os.Open(context.String("predictor"))
	if err != nil {
		PrintError(ERROR_LOADING_PREDICTOR, err)
		return
	}
	defer predictorFile.Close()

	gzipReader, err := gzip.NewReader(predictorFile)
	if err != nil {
		PrintError(ERROR_LOADING_PREDICTOR, err)
		return
	}
	defer gzipReader.Close()

	p, err := predictor.Load(gzipReader)
	if err != nil {
		PrintError(ERROR_LOADING_PREDICTOR, err)
		return
	}

	s := kompl.NewServer(context.String("port"), p)
	if err := s.Run(); err != nil {
		PrintError(ERROR_RUNNING_SERVER, err)
		return
	}
}
