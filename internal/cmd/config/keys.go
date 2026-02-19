package config

import (
	"fmt"
	"strconv"

	"git-sage/internal/config"
)

var configKeys = []string{"provider", "model", "instructions", "api_key", "token_limit"}

func getConfigValue(cfg *config.Config, key string) (any, bool) {
	switch key {
	case "provider":
		return cfg.Provider, true
	case "model":
		return cfg.Model, true
	case "instructions":
		return cfg.Instructions, true
	case "api_key":
		return cfg.APIKey, true
	case "token_limit":
		return cfg.TokenLimit, true
	default:
		return nil, false
	}
}

func setConfigValue(cfg *config.Config, key string, value string) error {
	switch key {
	case "provider":
		cfg.Provider = value
		return nil
	case "model":
		cfg.Model = value
		return nil
	case "instructions":
		cfg.Instructions = value
		return nil
	case "api_key":
		cfg.APIKey = value
		return nil
	case "token_limit":
		n, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("invalid number %q for key %q", value, key)
		}
		cfg.TokenLimit = n
		return nil
	default:
		return fmt.Errorf("unknown config key %q", key)
	}
}

func isEmpty(v any) bool {
	switch x := v.(type) {
	case string:
		return x == ""
	case int:
		return false
	default:
		return v == nil
	}
}

func formatValue(v any) string {
	switch x := v.(type) {
	case int:
		return strconv.Itoa(x)
	default:
		return fmt.Sprintf("%v", v)
	}
}
