package command

import (
	"os"

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
				Name:  "predictor,p",
				Value: "predictor.compl",
				Usage: "The path of a word predictor.",
			},

			cli.StringFlag{
				Name:  "port,n",
				Value: "8080",
				Usage: "The port number which a compl server uses.",
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

	predictor, err := compl.InflatePredictor(predictorFile)
	if err != nil {
		// TODO: Handle an error.
		return
	}

	candidateSeq := predictor.Predict([]string{"when", "kept"}, "a", 10)
	for _, candidate := range candidateSeq {
		println(candidate)
	}
	println("(predicted)")

	s := compl.NewServer(context.String("port"), predictor)
	if err := s.Run(); err != nil {
		// TODO: Handle an error.
		return
	}
}
