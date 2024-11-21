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


func TestValidateEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"john.doe@example.com", true},
		{"jane.doe123@sub.domain.org", true},
		{"invalid-email", false},
		{"@missingusername.com", false},
		{"missingdomain@.com", false},
		{"missingatsymbol.com", false},
		{"", false},
	}

	for _, test := range tests {
		result := ValidateEmail(test.input)
		if result != test.expected {
			t.Errorf("ValidateEmail(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}