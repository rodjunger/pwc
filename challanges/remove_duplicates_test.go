package pwc_test

import (
	"testing"

	pwc "github.com/rodjunger/pwc/challanges"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "Helo, Wrd!"},
		{"Hello", "Helo"},
		{"", ""},
		{"AAAAA", "A"},
		{"abcabc", "abc"},
		{"abababbaab", "ab"},
	}

	for _, test := range tests {
		result := pwc.RemoveDuplicateLetters(test.input)
		if result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Result: '%s'", test.input, test.expected, result)
		}
	}
}
