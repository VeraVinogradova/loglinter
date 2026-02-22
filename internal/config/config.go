package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	SensitiveWords []string          `json:"sensitive_words"`
	CustomPatterns map[string]string `json:"custom_patterns"`
	Fix            bool              `json:"auto_fix"`
	DisabledRules  []string          `json:"disabled_rules"`
}

func Load() (*Config, error) {
	cfg := defaultConfig()

	data, err := os.ReadFile(".loglinter.json")
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}

	var file Config
	if err := json.Unmarshal(data, &file); err != nil {
		return nil, err
	}

	return merge(cfg, &file), nil
}

func defaultConfig() *Config {
	return &Config{
		SensitiveWords: []string{
			"password", "secret", "token", "key",
			"credential", "auth", "private", "cert",
		},
		CustomPatterns: make(map[string]string),
		Fix:            false,
		DisabledRules:  []string{},
	}
}

func merge(dst, src *Config) *Config {
	if len(src.SensitiveWords) > 0 {
		seen := make(map[string]bool)
		for _, w := range dst.SensitiveWords {
			seen[w] = true
		}
		for _, w := range src.SensitiveWords {
			if !seen[w] {
				dst.SensitiveWords = append(dst.SensitiveWords, w)
			}
		}
	}

	for k, v := range src.CustomPatterns {
		dst.CustomPatterns[k] = v
	}

	if src.Fix {
		dst.Fix = src.Fix
	}

	if len(src.DisabledRules) > 0 {
		dst.DisabledRules = src.DisabledRules
	}

	return dst
}

func (c *Config) IsEnabled(rule string) bool {
	for _, r := range c.DisabledRules {
		if r == rule {
			return false
		}
	}
	return true
}
