package anthropic

import (
	"context"
	"strings"

	"github.com/iamtraction/sage/internal/llm"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type Client struct {
	client anthropic.Client
}

func New(apiKey string) (*Client, error) {
	client := anthropic.NewClient(option.WithAPIKey(apiKey))
	return &Client{client: client}, nil
}

func init() {
	llm.Register(llm.Anthropic, func(apiKey string) (llm.Provider, error) {
		return New(apiKey)
	})
}

const defaultModel = "claude-3-haiku"

func (c *Client) Generate(ctx context.Context, req llm.Request) (string, error) {
	if req.System == "" && req.User == "" {
		return "", nil
	}

	model := req.Model
	if model == "" {
		model = defaultModel
	}

	userContent := req.User
	if userContent == "" {
		userContent = "."
	}

	params := anthropic.MessageNewParams{
		MaxTokens: 4096,
		Messages:  []anthropic.MessageParam{anthropic.NewUserMessage(anthropic.NewTextBlock(userContent))},
		Model:     anthropic.Model(model),
	}
	if req.System != "" {
		params.System = []anthropic.TextBlockParam{{Text: req.System}}
	}

	message, err := c.client.Messages.New(ctx, params)
	if err != nil {
		return "", err
	}

	var text strings.Builder
	for _, block := range message.Content {
		if block.Text != "" {
			text.WriteString(block.Text)
		}
	}
	return text.String(), nil
}
