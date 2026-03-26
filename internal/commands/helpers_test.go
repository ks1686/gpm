package commands

import "testing"

func TestRedactValue(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		sensitive bool
		want      string
	}{
		{
			name:      "not sensitive, non-empty value",
			value:     "hello",
			sensitive: false,
			want:      "hello",
		},
		{
			name:      "not sensitive, empty value",
			value:     "",
			sensitive: false,
			want:      "",
		},
		{
			name:      "sensitive, non-empty value",
			value:     "secret",
			sensitive: true,
			want:      "[redacted]",
		},
		{
			name:      "sensitive, empty value",
			value:     "",
			sensitive: true,
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RedactValue(tt.value, tt.sensitive)
			if got != tt.want {
				t.Errorf("RedactValue(%q, %v) = %q, want %q", tt.value, tt.sensitive, got, tt.want)
			}
		})
	}
}
