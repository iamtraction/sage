package llm

import "errors"

type providerConstructor func(apiKey string) (Provider, error)

var registry = make(map[Name]providerConstructor)

func Register(name Name, fn providerConstructor) {
	registry[name] = fn
}

func New(name Name, apiKey string) (Provider, error) {
	if fn, ok := registry[name]; ok {
		return fn(apiKey)
	}
	return nil, errors.New("provider not implemented")
}
