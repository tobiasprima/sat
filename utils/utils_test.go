package utils

import (
	"testing"
)

func TestCapitalizeName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"john doe", "John Doe"},
		{"JANE DOE", "Jane Doe"},
		{"mary ann smith", "Mary Ann Smith"},
		{"", ""},
		{"  john   doe  ", "John Doe"},
	}

	for _, test := range tests {
		result := CapitalizeName(test.input)
		if result != test.expected {
			t.Errorf("CapitalizeName(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
