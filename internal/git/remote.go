package git

import (
	"context"
	"os/exec"
	"strings"
)

// GetRemotes returns the remote URLs for the repository.
func GetRemotes(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "remote", "-v")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
