package config

import (
	"fmt"

	"github.com/iamtraction/sage/internal/config"
	"github.com/iamtraction/sage/internal/logger"
)

func runShow() int {
	cfg, err := config.Load()
	if err != nil {
		return logger.Fatal("%v", err)
	}

	for _, key := range configKeys {
		v, ok := getConfigValue(cfg, key)
		if !ok {
			continue
		}
		fmt.Printf("%s = %s\n", key, formatValue(v))
	}
	return 0
}
