package git

import (
	"context"
	"os/exec"
)

// HasStagedChanges reports whether the working tree has staged changes.
func HasStagedChanges(ctx context.Context) (bool, error) {
	// exit 0 means no staged changes; exit 1 means staged changes exist.
	cmd := exec.CommandContext(ctx, "git", "diff", "--cached", "--quiet")
	err := cmd.Run()
	if err == nil {
		return false, nil
	}
	if cmd.ProcessState != nil && cmd.ProcessState.ExitCode() == 1 {
		return true, nil
	}
	return false, err
}
