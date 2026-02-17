package git

import (
	"context"
	"os/exec"
	"strings"
)

// IsGitRepo reports whether the current directory is inside a git working tree.
func IsGitRepo(ctx context.Context) (bool, error) {
	cmd := exec.CommandContext(ctx, "git", "rev-parse", "--is-inside-work-tree")
	out, err := cmd.Output()
	if err != nil {
		if cmd.ProcessState != nil && cmd.ProcessState.ExitCode() == 128 {
			return false, nil
		}
		return false, err
	}
	return strings.TrimSpace(string(out)) == "true", nil
}
