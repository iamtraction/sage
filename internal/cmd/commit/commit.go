package commit

import (
	"context"
	"fmt"

	"git-sage/internal/config"
	"git-sage/internal/git"
	"git-sage/internal/llm"
	_ "git-sage/internal/llm/anthropic"
	_ "git-sage/internal/llm/claude-code"
	_ "git-sage/internal/llm/google"
	_ "git-sage/internal/llm/openai"
	"git-sage/internal/logger"
	"git-sage/internal/prompts"
	"git-sage/internal/tokens"
)

const tokenLimit = 250_000

func Run(args []string) int {
	_ = args
	ctx := context.Background()

	// config

	cfg, err := config.Load()
	if err != nil {
		return logger.Fatal("load config: %v", err)
	}

	provider, err := llm.New(llm.Name(cfg.Provider), cfg.APIKey)
	if err != nil {
		return logger.Fatal("create provider: %v", err)
	}

	// context

	branch, err := git.GetCurrentBranch(ctx)
	if err != nil {
		return logger.Fatal("get current branch: %v", err)
	}

	hasStaged, err := git.HasStagedChanges(ctx)
	if err != nil {
		return logger.Fatal("has staged changes: %v", err)
	}
	if !hasStaged {
		logger.Info("no staged changes")
		return 0
	}

	nameStatus, err := git.GetNameStatus(ctx)
	if err != nil {
		return logger.Fatal("get name status: %v", err)
	}

	diff, err := git.GetStagedDiff(ctx, 3)
	if err != nil {
		return logger.Fatal("get staged diff: %v", err)
	}
	if tokens.Estimate(diff) > tokenLimit {
		diff, err = git.GetStagedDiff(ctx, 0)
		if err != nil {
			return logger.Fatal("get staged diff: %v", err)
		}
	}
	// TODO: implement diff truncation

	// prompts

	systemPrompt, err := prompts.Get("commit", nil)
	if err != nil {
		return logger.Fatal("get system prompt: %v", err)
	}

	userPrompt, err := prompts.Get("commit-metadata", map[string]string{
		"BranchName":       branch,
		"NameStatus":       git.FormatNameStatus(nameStatus),
		"StagedDiff":       string(diff),
		"UserInstructions": cfg.Instructions,
	})
	if err != nil {
		return logger.Fatal("get user prompt: %v", err)
	}

	// generation

	message, err := provider.Generate(ctx, llm.Request{
		Model:  cfg.Model,
		System: systemPrompt,
		User:   userPrompt,
	})
	if err != nil {
		return logger.Fatal("generate response: %v", err)
	}

	// commit

	out, err := git.Commit(ctx, message)
	if err != nil {
		return logger.Fatal("commit: %v", err)
	}

	fmt.Println(out)

	return 0
}
