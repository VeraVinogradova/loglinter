package testdata

import (
	"testing"

	"github.com/VeraVinogradova/loglinter/internal/config"
	"github.com/VeraVinogradova/loglinter/internal/rules"
)

func TestCheckSpecialChars(t *testing.T) {
	cfg := &config.Config{}

	tests := []struct {
		name string
		msg  string
		want bool
		fix  string
	}{
		{
			name: "empty_string",
			msg:  "",
			want: true,
			fix:  "",
		},
		{
			name: "valid_text",
			msg:  "hello world",
			want: true,
			fix:  "",
		},
		{
			name: "valid_single_exclamation",
			msg:  "hello!",
			want: true,
			fix:  "",
		},
		{
			name: "valid_single_question",
			msg:  "hello?",
			want: true,
			fix:  "",
		},
		{
			name: "valid_single_dot",
			msg:  "hello.",
			want: true,
			fix:  "",
		},
		{
			name: "invalid_double_exclamation",
			msg:  "hello!!",
			want: false,
			fix:  "hello",
		},
		{
			name: "invalid_triple_exclamation",
			msg:  "hello!!!",
			want: false,
			fix:  "hello",
		},
		{
			name: "invalid_multiple_dots",
			msg:  "hello...",
			want: false,
			fix:  "hello",
		},
		{
			name: "invalid_multiple_questions",
			msg:  "hello??",
			want: false,
			fix:  "hello",
		},
		{
			name: "invalid_mixed_punctuation",
			msg:  "hello!?",
			want: false,
			fix:  "hello",
		},
		{
			name: "invalid_emoji_only",
			msg:  "🚀",
			want: false,
			fix:  "",
		},
		{
			name: "invalid_emoji_with_text",
			msg:  "hello 🚀",
			want: false,
			fix:  "hello ",
		},
		{
			name: "invalid_multiple_emoji",
			msg:  "🚀🚀 hello",
			want: false,
			fix:  " hello",
		},
		{
			name: "invalid_emoji_and_punctuation",
			msg:  "hello!! 🚀",
			want: false,
			fix:  "hello ",
		},
		{
			name: "invalid_punctuation_at_start",
			msg:  "!!!hello",
			want: false,
			fix:  "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, fix := rules.CheckSpecialChars(tt.msg, cfg)
			if ok != tt.want {
				t.Errorf("CheckSpecialChars() ok = %v, want %v", ok, tt.want)
			}
			if fix != tt.fix {
				t.Errorf("CheckSpecialChars() fix = %q, want %q", fix, tt.fix)
			}
		})
	}
}
