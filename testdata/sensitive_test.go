package testdata

import (
	"testing"

	"github.com/VeraVinogradova/loglinter/internal/config"
	"github.com/VeraVinogradova/loglinter/internal/rules"
)

func TestCheckSensitive(t *testing.T) {
	cfg := &config.Config{
		SensitiveWords: []string{"password", "token", "secret", "key"},
		CustomPatterns: map[string]string{
			"email": `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`,
			"ip":    `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`,
		},
	}

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
			msg:  "user authenticated",
			want: true,
			fix:  "",
		},
		{
			name: "valid_similar_word",
			msg:  "passage",
			want: true,
			fix:  "",
		},
		{
			name: "invalid_password_exact",
			msg:  "password: 123",
			want: false,
			fix:  "",
		},
		{
			name: "invalid_password_case",
			msg:  "PASSWORD",
			want: false,
			fix:  "",
		},
		{
			name: "invalid_token",
			msg:  "token=abc123",
			want: false,
			fix:  "",
		},
		{
			name: "invalid_secret",
			msg:  "my secret",
			want: false,
			fix:  "",
		},
		{
			name: "invalid_api_key",
			msg:  "api_key=xyz",
			want: false,
			fix:  "",
		},
		{
			name: "invalid_multiple_words",
			msg:  "user password and token",
			want: false,
			fix:  "",
		},
		{
			name: "invalid_email_pattern",
			msg:  "user@example.com",
			want: false,
			fix:  "[REDACTED email]",
		},
		{
			name: "invalid_ip_pattern",
			msg:  "ip: 192.111.1.1",
			want: false,
			fix:  "[REDACTED ip]",
		},
		{
			name: "invalid_multiple_patterns",
			msg:  "email: user@example.com, ip: 10.0.0.1",
			want: false,
			fix:  "[REDACTED email]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, fix := rules.CheckSensitive(tt.msg, cfg)
			if ok != tt.want {
				t.Errorf("CheckSensitive() ok = %v, want %v", ok, tt.want)
			}
			if fix != tt.fix {
				t.Errorf("CheckSensitive() fix = %q, want %q", fix, tt.fix)
			}
		})
	}
}

func TestCheckSensitiveWithoutRules(t *testing.T) {
	cfg := &config.Config{
		SensitiveWords: []string{"password"},
		DisabledRules:  []string{"sensitive"},
	}

	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{
			name: "rule_disabled",
			msg:  "password: 123",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, _ := rules.CheckSensitive(tt.msg, cfg)
			if ok != tt.want {
				t.Errorf("CheckSensitive() = %v, want %v", ok, tt.want)
			}
		})
	}
}
