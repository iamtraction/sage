package git

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// GetRecentLog returns the most recent n commits as oneline entries.
func GetRecentLog(ctx context.Context, n int) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "log", "--oneline", fmt.Sprintf("-%d", n))
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// GetRecentTags returns the most recent n tags sorted by creation date.
func GetRecentTags(ctx context.Context, n int) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "tag", "--sort=-creatordate")
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
