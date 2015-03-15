package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/codegangsta/cli"
)

const (
	NAME         = "tokenize"
	VERSION      = "0.0.1"
	DESC         = "A stuplid tokenizer for English sentences."
	AUTHOR       = "Tomoya Kose (mitsuse)"
	AUTHOR_EMAIL = "tomoya@mitsuse.jp"

	BLANK_PATTERN  = `^\s*$`
	SYMBOL_PATTERN = `([~!@#\$%\^&\*\(\)\-_\+=\[\]\{\}\|\\;:"',<>\/\?])`
	DOT_PATTERN    = `([^A-Z\.])\.`
	SPACE_PATTERN  = `\s+`
	START_PATTERN  = `^\s+`
	END_PATTERN    = `\s+$`
	EOS_PATTERN    = `([\.!\?])\s+([A-Z"])`
)

var (
	blankRegexp  *regexp.Regexp
	symbolRegexp *regexp.Regexp
	dotRegexp    *regexp.Regexp
	spaceRegexp  *regexp.Regexp
	startRegexp  *regexp.Regexp
	endRegexp    *regexp.Regexp
	eosRegexp    *regexp.Regexp
)

func init() {
	blankRegexp = regexp.MustCompile(BLANK_PATTERN)
	symbolRegexp = regexp.MustCompile(SYMBOL_PATTERN)
	spaceRegexp = regexp.MustCompile(SPACE_PATTERN)
	dotRegexp = regexp.MustCompile(DOT_PATTERN)
	startRegexp = regexp.MustCompile(START_PATTERN)
	endRegexp = regexp.MustCompile(END_PATTERN)
	eosRegexp = regexp.MustCompile(EOS_PATTERN)
}

func main() {
	app := initApp()
	app.Run(os.Args)
}

func initApp() *cli.App {
	app := cli.NewApp()

	app.Name = NAME
	app.Version = VERSION
	app.Usage = DESC

	app.Author = AUTHOR
	app.Email = AUTHOR_EMAIL

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input,i",
			Value: "input.dump",
			Usage: "The path of input dump file.",
		},
	}
	app.Action = preproc

	return app
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", NAME, err)
}

func preproc(ctx *cli.Context) {
	dumpPath := ctx.String("input")
	dumpFile, err := os.Open(dumpPath)
	if err != nil {
		printError(err)
		return
	}
	defer dumpFile.Close()

	scanner := bufio.NewScanner(dumpFile)

	for scanner.Scan() {
		line := scanner.Text()

		if err := printProcessed(line); err != nil {
			printError(err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		printError(err)
		return
	}
}

func printProcessed(s string) error {
	if blankRegexp.MatchString(s) {
		return nil
	}

	processedString := symbolRegexp.ReplaceAllString(s, " $1 ")
	processedString = dotRegexp.ReplaceAllString(processedString, "$1 .")
	processedString = spaceRegexp.ReplaceAllString(processedString, " ")
	processedString = startRegexp.ReplaceAllString(processedString, "")
	processedString = endRegexp.ReplaceAllString(processedString, "")
	processedString = eosRegexp.ReplaceAllString(processedString, "$1\n$2")
	_, err := fmt.Println(processedString)

	return err
}
