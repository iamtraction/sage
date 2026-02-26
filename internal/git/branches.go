package git

import (
	"context"
	"os/exec"
	"strings"
)

// GetRecentBranches returns the most recent n branches sorted by commit date.
func GetRecentBranches(ctx context.Context, n int) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "branch", "--sort=-committerdate", "--format=%(refname:short)")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) > n {
		lines = lines[:n]
	}
	return strings.Join(lines, "\n"), nil
}
