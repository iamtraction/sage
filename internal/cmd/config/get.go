package config

import (
	"fmt"

	"git-sage/internal/config"
	"git-sage/internal/logger"
)

func runGet(key string) int {
	cfg, err := config.Load()
	if err != nil {
		return logger.Fatal("%v", err)
	}

	v, ok := getConfigValue(cfg, key)
	if !ok {
		return logger.Fatal("unknown config key %q", key)
	}
	if isEmpty(v) {
		return logger.Fatal("config key %q is not set", key)
	}
	fmt.Println(formatValue(v))
	return 0
}
