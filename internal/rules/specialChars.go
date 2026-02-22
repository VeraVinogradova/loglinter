package rules

import (
	"regexp"

	"github.com/VeraVinogradova/loglinter/internal/config"
)

var (
	emoji = regexp.MustCompile(`[\x{1F600}-\x{1F6FF}]|[\x{2600}-\x{26FF}]`)
	multi = regexp.MustCompile(`[!?.]{2,}`)
)

func CheckSpecialChars(msg string, cfg *config.Config) (bool, string) {
	if !cfg.IsEnabled("specialchars") {
		return true, ""
	}

	fixed := msg
	changed := false

	if emoji.MatchString(msg) {
		fixed = emoji.ReplaceAllString(fixed, "")
		changed = true
	}

	if multi.MatchString(fixed) {
		fixed = multi.ReplaceAllStringFunc(fixed, func(s string) string {
			return "" // возвращаем пустую строку вместо s[:1]
		})
		changed = true
	}

	if changed {
		return false, fixed
	}

	return true, ""
}
