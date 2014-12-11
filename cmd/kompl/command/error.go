package command

import (
	"fmt"
	"os"
)

func PrintError(message string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s: %s\n", Name, message, err)
}
