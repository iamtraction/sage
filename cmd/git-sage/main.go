package main

import (
	"context"
	"fmt"
	"os"

	"git-sage/internal/git"
	"git-sage/internal/tokens"
)

const tokenLimit = 100_000

func main() {
	ctx := context.Background()
	inside, err := git.IsGitRepo(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "git-sage: %v\n", err)
		os.Exit(1)
	}
	if !inside {
		fmt.Fprintln(os.Stderr, "git-sage: not a git repository")
		os.Exit(1)
	}
	hasStaged, err := git.HasStagedChanges(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "git-sage: %v\n", err)
		os.Exit(1)
	}
	if !hasStaged {
		fmt.Fprintln(os.Stderr, "git-sage: no staged changes")
		os.Exit(0)
	}
	diff, err := git.GetStagedDiff(ctx, 3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "git-sage: %v\n", err)
		os.Exit(1)
	}
	if tokens.Estimate(diff) > tokenLimit {
		diff, err = git.GetStagedDiff(ctx, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "git-sage: %v\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("%d tokens\n", tokens.Estimate(diff))
}
