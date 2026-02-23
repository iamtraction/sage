package config

import (
	"strings"

	"github.com/iamtraction/sage/internal/config"
	"github.com/iamtraction/sage/internal/logger"
)

func runSet(key string, valueParts []string) int {
	value := strings.Join(valueParts, " ")

	cfg, err := config.LoadOrEmpty()
	if err != nil {
		return logger.Fatal("%v", err)
	}

	if err := setConfigValue(cfg, key, value); err != nil {
		return logger.Fatal("%v", err)
	}

	if err := config.Save(cfg); err != nil {
		return logger.Fatal("%v", err)
	}
	return 0
}
