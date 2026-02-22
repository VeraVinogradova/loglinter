package testdata

import (
	"testing"

	"github.com/VeraVinogradova/loglinter/internal/config"
	"github.com/VeraVinogradova/loglinter/internal/rules"
)

func TestCheckLowercase(t *testing.T) {
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
			name: "valid_lowercase",
			msg:  "hello world",
			want: true,
			fix:  "",
		},
		{
			name: "valid_single_lowercase",
			msg:  "a",
			want: true,
			fix:  "",
		},
		{
			name: "invalid_uppercase_first",
			msg:  "Hello world",
			want: false,
			fix:  "hello world",
		},
		{
			name: "invalid_all_uppercase",
			msg:  "HELLO",
			want: false,
			fix:  "hELLO",
		},
		{
			name: "valid_starts_with_number",
			msg:  "1test",
			want: true,
			fix:  "",
		},
		{
			name: "valid_starts_with_symbol",
			msg:  "@test",
			want: true,
			fix:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, fix := rules.CheckLowercase(tt.msg, cfg)
			if ok != tt.want {
				t.Errorf("CheckLowercase() ok = %v, want %v", ok, tt.want)
			}
			if fix != tt.fix {
				t.Errorf("CheckLowercase() fix = %q, want %q", fix, tt.fix)
			}
		})
	}
}
