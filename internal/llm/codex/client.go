package codex

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"

	"git-sage/internal/llm"
)

type Client struct{}

func New(_ string) (*Client, error) {
	if _, err := exec.LookPath("codex"); err != nil {
		return nil, fmt.Errorf("codex CLI not found on PATH: %w", err)
	}
	return &Client{}, nil
}

func init() {
	llm.Register(llm.CodexCLI, func(apiKey string) (llm.Provider, error) {
		return New(apiKey)
	})
}

func (c *Client) Generate(ctx context.Context, req llm.Request) (string, error) {
	if req.System == "" && req.User == "" {
		return "", nil
	}

	// codex exec has no --system-prompt flag, so we prepend system instructions to the user prompt
	var prompt strings.Builder
	if req.System != "" {
		prompt.WriteString(req.System)
		prompt.WriteString("\n\n")
	}
	userPrompt := req.User
	if userPrompt == "" {
		userPrompt = "."
	}
	prompt.WriteString(userPrompt)

	args := []string{"exec", "--ephemeral"}

	if req.Model != "" {
		args = append(args, "--model", req.Model)
	}

	args = append(args, prompt.String())

	cmd := exec.CommandContext(ctx, "codex", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		errMsg := stderr.String()
		if errMsg == "" {
			errMsg = stdout.String()
		}
		return "", fmt.Errorf("codex CLI: %w: %s", err, strings.TrimSpace(errMsg))
	}

	return strings.TrimSpace(stdout.String()), nil
}
