package google

import (
	"context"

	"git-sage/internal/llm"

	"google.golang.org/genai"
)

type Client struct {
	client *genai.Client
}

func New(apiKey string) (*Client, error) {
	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}

func init() {
	llm.Register(llm.Google, func(apiKey string) (llm.Provider, error) {
		return New(apiKey)
	})
}

const defaultModel = "gemini-2.0-flash-lite"

func (c *Client) Generate(ctx context.Context, req llm.Request) (string, error) {
	if req.System == "" && req.User == "" {
		return "", nil
	}

	model := req.Model
	if model == "" {
		model = defaultModel
	}

	config := &genai.GenerateContentConfig{}
	if req.System != "" {
		config.SystemInstruction = &genai.Content{
			Parts: []*genai.Part{{Text: req.System}},
		}
	}

	userContent := req.User
	if userContent == "" {
		userContent = "."
	}
	parts := []*genai.Part{{Text: userContent}}
	result, err := c.client.Models.GenerateContent(ctx, model, []*genai.Content{{Parts: parts}}, config)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}
