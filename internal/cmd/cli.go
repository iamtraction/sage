package cmd

import (
	"context"

	"github.com/iamtraction/sage/internal/cmd/commit"
	"github.com/iamtraction/sage/internal/cmd/config"
	"github.com/iamtraction/sage/internal/git"
	"github.com/iamtraction/sage/internal/logger"
	"github.com/iamtraction/sage/internal/version"
)

func Run(args []string) int {
	ctx := context.Background()

	if len(args) < 1 {
		return runDefault()
	}

	if args[0] == "version" {
		logger.Info(version.Version)
		return 0
	}

	inside, err := git.IsGitRepo(ctx)
	if err != nil {
		return logger.Fatal("%v", err)
	}
	if !inside {
		return logger.Fatal("not a git repository")
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
