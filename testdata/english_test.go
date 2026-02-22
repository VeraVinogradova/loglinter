package testdata

import (
	"testing"

	"github.com/VeraVinogradova/loglinter/internal/config"
	"github.com/VeraVinogradova/loglinter/internal/rules"
)

func TestCheckEnglish(t *testing.T) {
	cfg := &config.Config{}

	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{
			name: "empty_string",
			msg:  "",
			want: true,
		},
		{
			name: "valid_english",
			msg:  "hello world",
			want: true,
		},
		{
			name: "valid_with_numbers",
			msg:  "hello123",
			want: true,
		},
		{
			name: "valid_with_symbols",
			msg:  "hello-world",
			want: true,
		},
		{
			name: "russian",
			msg:  "привет мир",
			want: false,
		},
		{
			name: "hindi",
			msg:  "हैलो वर्ल्ड",
			want: false,
		},
		{
			name: "arab",
			msg:  "مرحبا",
			want: false,
		},
		{
			name: "chinese",
			msg:  "你好",
			want: false,
		},
		{
			name: "mixed_english_russian",
			msg:  "hello привет",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, _ := rules.CheckEnglish(tt.msg, cfg)
			if ok != tt.want {
				t.Errorf("CheckEnglish() = %v, want %v", ok, tt.want)
			}
		})
	}
}
