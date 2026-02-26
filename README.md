# sage

AI-powered Git intelligence assistant. Use natural language to execute git and gh commands. Or, simply generate [Conventional Commit](https://www.conventionalcommits.org/) messages from staged changes.

## Install

### Linux / macOS

```sh
curl -sSL https://raw.github.com/iamtraction/sage/main/install.sh | sh
```

Installs to `~/.local/bin/sage`. Make sure `~/.local/bin` is in your `PATH`.

### Windows (PowerShell)

```powershell
iex "& { $(iwr -useb https://raw.github.com/iamtraction/sage/main/install.ps1) }"
```

Installs to `%LOCALAPPDATA%\bin\sage.exe`. Make sure `%LOCALAPPDATA%\bin` is in your `PATH`.

### From Source

```sh
go install github.com/iamtraction/sage@latest
```

## Quick Start

```sh
# 1. Set your provider
sage config provider anthropic

# 2. Set your API key
sage config api_key sk-ant-...

# 3. Stage changes and commit
git add .
sage commit

# or use natural language to run git commands
sage exec show me the last 5 commits
```

## Configuration

Configuration is stored in `~/.config/sage/config.json` (Linux/macOS) or `%LOCALAPPDATA%\sage\config.json` (Windows).

```sh
sage config                        # show all config
sage config <key>                  # get a value
sage config <key> <value>          # set a value
```

### Config Keys

| Key | Description |
|-----|-------------|
| `provider` | LLM provider to use (see [Providers](#providers)) |
| `model` | Model override (optional — each provider has a sensible default) |
| `api_key` | API key for SDK-based providers |
| `instructions` | Custom instructions to influence commit message style |
| `auto_execute` | Auto-execute non-destructive commands from `sage exec` (`true`/`false`) |

### Custom Instructions

You can add custom instructions to guide the commit message style:

```sh
sage config instructions "always use lowercase, keep subject under 50 chars"
```

## Providers

These providers call the LLM API directly. Set your `api_key` in config.

| Provider | Config Value | Default Model | API Key |
|----------|-------------|---------------|---------|
| Anthropic | `anthropic` | `claude-3-haiku` | [console.anthropic.com](https://console.anthropic.com/) |
| OpenAI | `openai` | `gpt-5-nano` | [platform.openai.com](https://platform.openai.com/) |
| Google Gemini | `google` | `gemini-2.0-flash-lite` | [aistudio.google.com](https://aistudio.google.com/) |

```sh
sage config provider anthropic
sage config api_key sk-ant-...
```

### CLI-based

These providers wrap an installed CLI tool. They use the CLI's own authentication — no `api_key` required.

| Provider | Config Value | CLI Required | Install |
|----------|-------------|-------------|---------|
| Claude Code | `claude-code` | `claude` | [claude.ai/download](https://claude.ai/download) |
| Codex | `codex` | `codex` | [github.com/openai/codex](https://github.com/openai/codex) |
| Gemini CLI | `gemini-cli` | `gemini` | `npm i -g @google/gemini-cli` |

```sh
sage config provider claude-code
# no api_key needed — just make sure `claude` is on your PATH
```

### Overriding the Model

Each provider has a default model, but you can override it:

```sh
sage config model claude-sonnet-4-6    # for anthropic
sage config model gpt-5                # for openai
sage config model gemini-2.5-flash     # for google
sage config model sonnet               # for claude-code (uses CLI aliases)
```

## Usage

### Commit

```sh
# stage your changes
git add -A

# generate commit message and commit
sage commit
```

sage analyzes your staged diff, generates a Conventional Commit message, and commits the changes for you.

### Exec

```sh
# generate and execute git/gh commands from natural language
sage exec list branches merged into main
sage exec show commits by user@example.com this week
sage exec create a tag v1.2.0 on HEAD

# skip confirmation
sage exec -y show recent tags

# auto-execute non-destructive commands
sage config auto_execute true
```

sage generates the appropriate git or gh command, shows it with a description, and asks for confirmation before executing. Destructive commands (force push, branch deletion, history rewriting) always require confirmation, even with `auto_execute` enabled.

## License

MIT
