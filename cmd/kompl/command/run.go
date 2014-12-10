package command

import (
	"compress/gzip"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/kompl"
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
				Value: "8080",
				Usage: "The port number which a kompl server uses.",
			},
		},
	}

	return command
}

func runAction(context *cli.Context) {
	// TODO: Start a seal server.
	predictorFile, err := os.Open(context.String("predictor"))
	if err != nil {
		// TODO: Handle an error.
		return
	}
	defer predictorFile.Close()

	gzipReader, err := gzip.NewReader(predictorFile)
	if err != nil {
		// TODO: Handle an error.
		return
	}
	defer gzipReader.Close()

	predictor, err := kompl.InflatePredictor(gzipReader)
	if err != nil {
		// TODO: Handle an error.
		return
	}

	s := kompl.NewServer(context.String("port"), predictor)
	if err := s.Run(); err != nil {
		// TODO: Handle an error.
		return
	}
}
