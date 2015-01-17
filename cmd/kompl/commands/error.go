package commands

import (
	"fmt"
	"os"
)

func PrintError(message string, err error) {
	if err == nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", Name, message)
	} else {
		fmt.Fprintf(os.Stderr, "%s: %s: %s\n", Name, message, err)
	}
}
