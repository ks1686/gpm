package commands

import (
	"strings"
	"testing"
)

func TestKnownManagerList(t *testing.T) {
	// The expected output depends on the keys of schema.KnownManagers being joined and sorted.
	// Since schema.KnownManagers is defined as:
	// "apt", "dnf", "pacman", "paru", "yay", "flatpak", "snap", "brew", "macports", "linuxbrew"
	expected := "apt, brew, dnf, flatpak, linuxbrew, macports, pacman, paru, snap, yay"

	result := KnownManagerList()
	if result != expected {
		t.Errorf("KnownManagerList() = %q, want %q", result, expected)
	}

	// Double check that it contains at least one known manager to be safe against schema changes
	if !strings.Contains(result, "apt") {
		t.Errorf("KnownManagerList() output does not contain expected manager %q", "apt")
	}
}

func TestRedactValue(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		sensitive bool
		want      string
	}{
		{"empty string, not sensitive", "", false, ""},
		{"empty string, sensitive", "", true, ""},
		{"value, not sensitive", "mysecret", false, "mysecret"},
		{"value, sensitive", "mysecret", true, "[redacted]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RedactValue(tt.value, tt.sensitive); got != tt.want {
				t.Errorf("RedactValue() = %q, want %q", got, tt.want)
			}
		})
	}
}
