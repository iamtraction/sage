package main

import (
	"os"

	"github.com/iamtraction/sage/internal/cmd"
)

func main() {
	os.Exit(cmd.Run(os.Args[1:]))
}
