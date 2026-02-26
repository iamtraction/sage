package codex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/iamtraction/sage/internal/llm"
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

	args := []string{"exec", "--ephemeral"}

	if req.System != "" {
		args = append(args, "-c", "developer_instructions="+req.System)
	}

	if req.Model != "" {
		args = append(args, "--model", req.Model)
	}

	if req.OutputSchema != nil {
		schemaFile, err := os.CreateTemp("", "sage-schema-*.json")
		if err != nil {
			return "", fmt.Errorf("create schema temp file: %w", err)
		}
		defer os.Remove(schemaFile.Name())

		if err := json.NewEncoder(schemaFile).Encode(req.OutputSchema); err != nil {
			schemaFile.Close()
			return "", fmt.Errorf("write schema: %w", err)
		}
		schemaFile.Close()

		args = append(args, "--output-schema", schemaFile.Name())
	}

	userPrompt := req.User
	if userPrompt == "" {
		userPrompt = "."
	}

	args = append(args, "-")

	cmd := exec.CommandContext(ctx, "codex", args...)
	cmd.Stdin = strings.NewReader(userPrompt)

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
