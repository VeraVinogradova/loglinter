package rules

import (
	"unicode"

	"github.com/VeraVinogradova/loglinter/internal/config"
)

var others = []*unicode.RangeTable{
	unicode.Cyrillic,
	unicode.Greek,
	unicode.Arabic,
	unicode.Hebrew,
	unicode.Han,
	unicode.Thai,
	unicode.Devanagari,
}

func CheckEnglish(msg string, cfg *config.Config) (bool, string) {
	if !cfg.IsEnabled("english") {
		return true, ""
	}

	for _, r := range msg {
		for _, t := range others {
			if unicode.Is(t, r) {
				return false, ""
			}
		}
	}
	return true, ""
}
