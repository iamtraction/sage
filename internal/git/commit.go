package git

import (
	"context"
	"os/exec"
	"strings"
)

// Commit commits the changes with the given message.
func Commit(ctx context.Context, message string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "commit", "-m", message)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
