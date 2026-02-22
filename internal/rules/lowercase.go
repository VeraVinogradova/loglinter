package rules

import (
	"strings"
	"unicode"

	"github.com/VeraVinogradova/loglinter/internal/config"
)

func CheckLowercase(msg string, cfg *config.Config) (bool, string) {
	if !cfg.IsEnabled("lowercase") {
		return true, ""
	}

	if len(msg) == 0 {
		return true, ""
	}

	if unicode.IsUpper(rune(msg[0])) {
		fixed := strings.ToLower(string(msg[0])) + msg[1:]
		return false, fixed
	}

	return true, ""
}
