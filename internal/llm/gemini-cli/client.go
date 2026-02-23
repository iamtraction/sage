package geminicli

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
	if _, err := exec.LookPath("gemini"); err != nil {
		return nil, fmt.Errorf("gemini CLI not found on PATH: %w", err)
	}
	return &Client{}, nil
}

func init() {
	llm.Register(llm.GeminiCLI, func(apiKey string) (llm.Provider, error) {
		return New(apiKey)
	})
}

func (c *Client) Generate(ctx context.Context, req llm.Request) (string, error) {
	if req.System == "" && req.User == "" {
		return "", nil
	}

	prompt := req.System
	if prompt == "" {
		prompt = "."
	}

	args := []string{"-p", prompt}

	if req.Model != "" {
		args = append(args, "-m", req.Model)
	}

	cmd := exec.CommandContext(ctx, "gemini", args...)

	if req.User != "" {
		cmd.Stdin = strings.NewReader(req.User)
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		errMsg := stderr.String()
		if errMsg == "" {
			errMsg = stdout.String()
		}
		return "", fmt.Errorf("gemini CLI: %w: %s", err, strings.TrimSpace(errMsg))
	}

	return strings.TrimSpace(stdout.String()), nil
}
