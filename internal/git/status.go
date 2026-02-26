package git

import (
	"bufio"
	"context"
	"os/exec"
	"strings"
)

// FileChange represents a staged file change
type FileChange struct {
	Status  string // A, M, D, R, or C
	Path    string
	OldPath string // old path for R and C
}

// GetNameStatus returns staged file changes and parses A, M, D, R, and C lines.
func GetNameStatus(ctx context.Context) ([]FileChange, error) {
	cmd := exec.CommandContext(ctx, "git", "diff", "--staged", "--name-status")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var changes []FileChange
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Split(line, "\t")
		switch len(fields) {
		case 2:
			changes = append(changes, FileChange{Status: fields[0], Path: fields[1]})
		case 3:
			changes = append(changes, FileChange{
				Status:  fields[0],
				OldPath: fields[1],
				Path:    fields[2],
			})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return changes, nil
}

// GetStatus returns the output of git status.
func GetStatus(ctx context.Context, short bool) (string, error) {
	args := []string{"status"}
	if short {
		args = append(args, "--short")
	}
	cmd := exec.CommandContext(ctx, "git", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// FormatNameStatus formats the name status for the prompt.
func FormatNameStatus(changes []FileChange) string {
	var sb strings.Builder
	for i, change := range changes {
		sb.WriteString(change.Status)
		sb.WriteString("\t")
		sb.WriteString(change.Path)
		if i != len(changes)-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
