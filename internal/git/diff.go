package git

import (
	"context"
	"fmt"
	"os/exec"
)

// GetDiff returns the diff with the given unified context line count.
func GetDiff(ctx context.Context, unified int) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "git", "diff", fmt.Sprintf("--unified=%d", unified))
	return cmd.Output()
}

// GetStagedDiff returns the staged diff with the given unified context line count.
func GetStagedDiff(ctx context.Context, unified int) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "git", "diff", "--staged", fmt.Sprintf("--unified=%d", unified))
	return cmd.Output()
}
