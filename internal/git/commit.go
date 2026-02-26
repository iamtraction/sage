package git

import (
	"context"
	"os/exec"
	"strings"
)

// Commit commits the changes with the given message.
func Commit(ctx context.Context, message string) (string, error) {
	message = message + "\n\nCo-Authored-By: sage <264273354+sage-cli@users.noreply.github.com>"
	cmd := exec.CommandContext(ctx, "git", "commit", "-s", "-m", message)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
