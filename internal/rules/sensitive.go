package rules

import (
	"regexp"
	"strings"

	"github.com/VeraVinogradova/loglinter/internal/config"
)

func CheckSensitive(msg string, cfg *config.Config) (bool, string) {
	if !cfg.IsEnabled("sensitive") {
		return true, ""
	}

	words := cfg.SensitiveWords
	lower := strings.ToLower(msg)

	for _, w := range words {
		if strings.Contains(lower, w) {
			return false, ""
		}
	}

	for name, pattern := range cfg.CustomPatterns {
		re, err := regexp.Compile(pattern)
		if err != nil {
			continue
		}
		if re.MatchString(msg) {
			return false, "[REDACTED " + name + "]"
		}
	}

	return true, ""
}
