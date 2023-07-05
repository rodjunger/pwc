package pwc_test

import (
	"testing"

	pwc "github.com/rodjunger/pwc/challanges"
)

func TestReverseWordOrder(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World! OpenAI is amazing.", "amazing. is OpenAI World! Hello,"},
		{"The quick brown fox", "fox brown quick The"},
		{"oneword", "oneword"},
		{"two   words", "words   two"},
	}

	for _, test := range tests {
		result := pwc.ReverseWordOrder(test.input)
		if result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Result: '%s'", test.input, test.expected, result)
		}
	}
}
