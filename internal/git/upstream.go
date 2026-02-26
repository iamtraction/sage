package git

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// GetUpstreamStatus returns the upstream tracking branch and ahead/behind counts.
func GetUpstreamStatus(ctx context.Context) (string, error) {
	// get upstream branch name
	cmd := exec.CommandContext(ctx, "git", "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}")
	out, err := cmd.Output()
	if err != nil {
		return "no upstream tracking branch", nil
	}
	upstream := strings.TrimSpace(string(out))

	// get ahead/behind counts
	cmd = exec.CommandContext(ctx, "git", "rev-list", "--left-right", "--count", "HEAD...@{u}")
	out, err = cmd.Output()
	if err != nil {
		return upstream, nil
	}
	parts := strings.Fields(strings.TrimSpace(string(out)))
	if len(parts) == 2 {
		return fmt.Sprintf("%s (ahead %s, behind %s)", upstream, parts[0], parts[1]), nil
	}
	return upstream, nil
}
