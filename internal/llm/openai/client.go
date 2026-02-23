package openai

import (
	"context"

	"github.com/iamtraction/sage/internal/llm"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type Client struct {
	client openai.Client
}

func New(apiKey string) (*Client, error) {
	client := openai.NewClient(option.WithAPIKey(apiKey))
	return &Client{client: client}, nil
}

func init() {
	llm.Register(llm.OpenAI, func(apiKey string) (llm.Provider, error) {
		return New(apiKey)
	})
}

const defaultModel = "gpt-5-nano"

func (c *Client) Generate(ctx context.Context, req llm.Request) (string, error) {
	if req.System == "" && req.User == "" {
		return "", nil
	}

	messages := make([]openai.ChatCompletionMessageParamUnion, 0, 2)
	if req.System != "" {
		messages = append(messages, openai.SystemMessage(req.System))
	}
	if req.User != "" {
		messages = append(messages, openai.UserMessage(req.User))
	}

	model := req.Model
	if model == "" {
		model = defaultModel
	}

	params := openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    openai.ChatModel(model),
	}

	completion, err := c.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}

	if len(completion.Choices) == 0 {
		return "", nil
	}
	return completion.Choices[0].Message.Content, nil
}
