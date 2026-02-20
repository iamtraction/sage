package llm

import "context"

type Provider interface {
	Generate(ctx context.Context, req Request) (string, error)
}
