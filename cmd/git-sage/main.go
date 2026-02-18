package main

import (
	"os"

	"git-sage/internal/cmd"
)

func main() {
	os.Exit(cmd.Run(os.Args[1:]))
}
