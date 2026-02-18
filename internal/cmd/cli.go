package cmd

import (
	"git-sage/internal/logger"
)

func Run(args []string) int {
	if len(args) < 1 {
		return runDefault()
	}
	switch args[0] {
	default:
		return logger.Fatal("unknown command %q", args[0])
	}
}
