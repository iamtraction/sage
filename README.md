# git-sage

AI-powered commit message generator. Stage your changes, run `git-sage commit`, and get a [Conventional Commit](https://www.conventionalcommits.org/) message — no thinking required.

## Install

### Linux / macOS

```sh
curl -sSL https://raw.github.com/iamtraction/sage/main/install.sh | sh
```

Installs to `~/.local/bin/git-sage`. Make sure `~/.local/bin` is in your `PATH`.

### Windows (PowerShell)

```powershell
iex (iwr https://raw.github.com/iamtraction/sage/main/install.ps1).Content
```

Installs to `%LOCALAPPDATA%\bin\git-sage.exe`. Make sure `%LOCALAPPDATA%\bin` is in your `PATH`.

### From Source

```sh
go install git-sage@latest
```

## Quick Start

```sh
# 1. Set your provider
git-sage config provider anthropic

# 2. Set your API key
git-sage config api_key sk-ant-...

# 3. Stage changes and commit
git add .
git-sage commit
```

## Configuration

Configuration is stored in `~/.config/git-sage/config.json` (Linux/macOS) or `%LOCALAPPDATA%\git-sage\config.json` (Windows).

```sh
git-sage config                        # show all config
git-sage config <key>                  # get a value
git-sage config <key> <value>          # set a value
```

### Config Keys

| Key | Description |
|-----|-------------|
| `provider` | LLM provider to use (see [Providers](#providers)) |
| `model` | Model override (optional — each provider has a sensible default) |
| `api_key` | API key for SDK-based providers |
| `instructions` | Custom instructions to influence commit message style |

### Custom Instructions

You can add custom instructions to guide the commit message style:

```sh
git-sage config instructions "always use lowercase, keep subject under 50 chars"
```

## Providers

These providers call the LLM API directly. Set your `api_key` in config.

| Provider | Config Value | Default Model | API Key |
|----------|-------------|---------------|---------|
| Anthropic | `anthropic` | `claude-3-haiku` | [console.anthropic.com](https://console.anthropic.com/) |
| OpenAI | `openai` | `gpt-5-nano` | [platform.openai.com](https://platform.openai.com/) |
| Google Gemini | `google` | `gemini-2.0-flash-lite` | [aistudio.google.com](https://aistudio.google.com/) |

```sh
git-sage config provider anthropic
git-sage config api_key sk-ant-...
```

### CLI-based

These providers wrap an installed CLI tool. They use the CLI's own authentication — no `api_key` required.

| Provider | Config Value | CLI Required | Install |
|----------|-------------|-------------|---------|
| Claude Code | `claude-code` | `claude` | [claude.ai/download](https://claude.ai/download) |
| Codex | `codex` | `codex` | [github.com/openai/codex](https://github.com/openai/codex) |
| Gemini CLI | `gemini-cli` | `gemini` | `npm i -g @google/gemini-cli` |

```sh
git-sage config provider claude-code
# no api_key needed — just make sure `claude` is on your PATH
```

### Overriding the Model

Each provider has a default model, but you can override it:

```sh
git-sage config model claude-sonnet-4-6    # for anthropic
git-sage config model gpt-5                # for openai
git-sage config model gemini-2.5-flash     # for google
git-sage config model sonnet               # for claude-code (uses CLI aliases)
```

To reset to the provider's default, clear the model:

```sh
git-sage config model ""
```

## Usage

```sh
# stage your changes
git add -A

# generate commit message and commit
git-sage commit
```

That's it. git-sage will analyze your staged diff, generate a Conventional Commit message, and run `git commit` for you.

## License

MIT
