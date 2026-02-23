package claudecode

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/iamtraction/sage/internal/llm"
)

type Client struct{}

func New(_ string) (*Client, error) {
	if _, err := exec.LookPath("claude"); err != nil {
		return nil, fmt.Errorf("claude CLI not found on PATH: %w", err)
	}
	return &Client{}, nil
}

func init() {
	llm.Register(llm.ClaudeCode, func(apiKey string) (llm.Provider, error) {
		return New(apiKey)
	})
}

func (c *Client) Generate(ctx context.Context, req llm.Request) (string, error) {
	if req.System == "" && req.User == "" {
		return "", nil
	}

	args := []string{"-p"}

	if req.System != "" {
		args = append(args, "--system-prompt", req.System)
	}

	if req.Model != "" {
		args = append(args, "--model", req.Model)
	}

	userPrompt := req.User
	if userPrompt == "" {
		userPrompt = "."
	}

	cmd := exec.CommandContext(ctx, "claude", args...)
	cmd.Stdin = strings.NewReader(userPrompt)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		// claude may write errors to stdout or stderr
		errMsg := stderr.String()
		if errMsg == "" {
			errMsg = stdout.String()
		}
		return "", fmt.Errorf("claude CLI: %w: %s", err, strings.TrimSpace(errMsg))
	}

	return strings.TrimSpace(stdout.String()), nil
}
