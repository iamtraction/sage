package main

import (
	"context"
	"fmt"
	"os"

	"git-sage/internal/git"
)

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
	fmt.Println("git-sage: run 'git sage help' for more information")
}
