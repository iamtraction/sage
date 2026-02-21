package cmd

import (
	"git-sage/internal/cmd/commit"
	"git-sage/internal/cmd/config"
	"git-sage/internal/logger"
)

func Run(args []string) int {
	if len(args) < 1 {
		return runDefault()
	}
	switch args[0] {
	case "commit":
		return commit.Run(args[1:])
	case "config":
		return config.Run(args[1:])
	default:
		return logger.Fatal("unknown command %q", args[0])
	}
}
