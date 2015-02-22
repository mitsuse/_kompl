package commands

import (
	"fmt"
	"os"
)

func PrintError(message string, err error) {
	if err == nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", NAME, message)
	} else {
		fmt.Fprintf(os.Stderr, "%s: %s: %s\n", NAME, message, err)
	}
}
