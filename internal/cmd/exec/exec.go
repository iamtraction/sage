package exec

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	osexec "os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/iamtraction/sage/internal/config"
	"github.com/iamtraction/sage/internal/git"
	"github.com/iamtraction/sage/internal/llm"
	_ "github.com/iamtraction/sage/internal/llm/anthropic"
	_ "github.com/iamtraction/sage/internal/llm/claude-code"
	_ "github.com/iamtraction/sage/internal/llm/codex"
	_ "github.com/iamtraction/sage/internal/llm/gemini-cli"
	_ "github.com/iamtraction/sage/internal/llm/google"
	_ "github.com/iamtraction/sage/internal/llm/openai"
	"github.com/iamtraction/sage/internal/logger"
	"github.com/iamtraction/sage/internal/prompts"
)

type execResponse struct {
	Command     string `json:"command"`
	Destructive bool   `json:"destructive"`
	Description string `json:"description"`
}

var execResponseSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"command": map[string]any{
			"type":        "string",
			"description": "The full git or gh command to run",
		},
		"destructive": map[string]any{
			"type":        "boolean",
			"description": "Whether the command modifies history, deletes data, or is hard to reverse",
		},
		"description": map[string]any{
			"type":        "string",
			"description": "Brief human-readable explanation of what the command does",
		},
	},
	"required":             []string{"command", "destructive", "description"},
	"additionalProperties": false,
}

func Run(args []string) int {
	ctx := context.Background()

	// parse flags
	autoYes := false
	var promptArgs []string
	for _, arg := range args {
		if strings.EqualFold(arg, "-y") || strings.EqualFold(arg, "--yes") {
			autoYes = true
		} else {
			promptArgs = append(promptArgs, arg)
		}
	}

	if len(promptArgs) == 0 {
		fmt.Println("Usage: sage exec [-y] <prompt>")
		fmt.Println()
		fmt.Println("Generate and execute git/gh commands from natural language.")
		fmt.Println()
		fmt.Println("Flags:")
		fmt.Println("  -y, --yes    execute without confirmation")
		return 1
	}

	request := strings.Join(promptArgs, " ")

	// config

	cfg, err := config.Load()
	if err != nil {
		return logger.Fatal("load config: %v", err)
	}

	provider, err := llm.New(llm.Name(cfg.Provider), cfg.APIKey)
	if err != nil {
		return logger.Fatal("create provider: %v", err)
	}

	// gather context in parallel

	type result struct {
		value string
		err   error
	}

	var (
		branch, upstream, remotes, recentBranches, status, recentLog, recentTags result
		wg                                                                       sync.WaitGroup
	)

	wg.Add(7)

	go func() {
		defer wg.Done()
		branch.value, branch.err = git.GetCurrentBranch(ctx)
	}()
	go func() {
		defer wg.Done()
		upstream.value, upstream.err = git.GetUpstreamStatus(ctx)
	}()
	go func() {
		defer wg.Done()
		remotes.value, remotes.err = git.GetRemotes(ctx)
	}()
	go func() {
		defer wg.Done()
		recentBranches.value, recentBranches.err = git.GetRecentBranches(ctx, 10)
	}()
	go func() {
		defer wg.Done()
		status.value, status.err = git.GetStatus(ctx, true)
	}()
	go func() {
		defer wg.Done()
		recentLog.value, recentLog.err = git.GetRecentLog(ctx, 15)
	}()
	go func() {
		defer wg.Done()
		recentTags.value, recentTags.err = git.GetRecentTags(ctx, 5)
	}()

	wg.Wait()

	if branch.err != nil {
		return logger.Fatal("get current branch: %v", branch.err)
	}

	// prompts

	systemPrompt, err := prompts.Get("exec", nil)
	if err != nil {
		return logger.Fatal("get system prompt: %v", err)
	}

	shell := "sh"
	if runtime.GOOS == "windows" {
		shell = "cmd"
	}

	userPrompt, err := prompts.Get("exec-metadata", map[string]string{
		"OS":             runtime.GOOS,
		"Shell":          shell,
		"BranchName":     branch.value,
		"Upstream":       upstream.value,
		"Remotes":        remotes.value,
		"RecentBranches": recentBranches.value,
		"Status":         status.value,
		"RecentLog":      recentLog.value,
		"RecentTags":     recentTags.value,
		"Request":        request,
	})
	if err != nil {
		return logger.Fatal("get user prompt: %v", err)
	}

	// generation

	raw, err := provider.Generate(ctx, llm.Request{
		Model:        cfg.Model,
		System:       systemPrompt,
		User:         userPrompt,
		OutputSchema: execResponseSchema,
	})
	if err != nil {
		return logger.Fatal("generate response: %v", err)
	}

	raw = extractJSON(raw)

	var resp execResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		return logger.Fatal("parse response: %v\nraw: %s", err, raw)
	}

	// validate command
	if !strings.HasPrefix(resp.Command, "git ") && !strings.HasPrefix(resp.Command, "gh ") {
		return logger.Fatal("generated command must start with 'git' or 'gh': %s", resp.Command)
	}

	// display
	fmt.Printf("  %s\n", resp.Command)
	fmt.Printf("  %s\n", resp.Description)

	// check gh availability
	if strings.HasPrefix(resp.Command, "gh ") {
		if _, err := osexec.LookPath("gh"); err != nil {
			fmt.Println()
			logger.Info("gh CLI is required but not installed. Install it from https://cli.github.com")
			return 1
		}
	}

	// confirmation logic
	if !autoYes && !(cfg.AutoExecute && !resp.Destructive) {
		fmt.Print("Execute? [y/n]: ")
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))
		if answer != "y" && answer != "yes" {
			return 0
		}
	}

	// execute
	fmt.Println()
	var cmd *osexec.Cmd
	if runtime.GOOS == "windows" {
		cmd = osexec.CommandContext(ctx, "cmd", "/c", resp.Command)
	} else {
		cmd = osexec.CommandContext(ctx, "sh", "-c", resp.Command)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return logger.Fatal("execute: %v", err)
	}

	return 0
}

// extractJSON finds the first top-level JSON object in s by matching braces.
func extractJSON(s string) string {
	start := strings.Index(s, "{")
	if start == -1 {
		return s
	}
	depth := 0
	inString := false
	escaped := false
	for i := start; i < len(s); i++ {
		c := s[i]
		if escaped {
			escaped = false
			continue
		}
		if c == '\\' && inString {
			escaped = true
			continue
		}
		if c == '"' {
			inString = !inString
			continue
		}
		if inString {
			continue
		}
		switch c {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return s[start : i+1]
			}
		}
	}
	return s[start:]
}
