package cmd

import (
	"context"
	"fmt"

	"git-sage/internal/git"
	"git-sage/internal/logger"
	"git-sage/internal/tokens"
)

const tokenLimit = 100_000

func runDefault() int {
	ctx := context.Background()

	inside, err := git.IsGitRepo(ctx)
	if err != nil {
		return logger.Fatal("%v", err)
	}
	if !inside {
		return logger.Fatal("not a git repository")
	}

	hasStaged, err := git.HasStagedChanges(ctx)
	if err != nil {
		return logger.Fatal("%v", err)
	}
	if !hasStaged {
		logger.Info("no staged changes")
		return 0
	}

	diff, err := git.GetStagedDiff(ctx, 3)
	if err != nil {
		return logger.Fatal("%v", err)
	}
	if tokens.Estimate(diff) > tokenLimit {
		diff, err = git.GetStagedDiff(ctx, 0)
		if err != nil {
			return logger.Fatal("%v", err)
		}
	}
	fmt.Printf("%d tokens\n", tokens.Estimate(diff))

	return 0
}
