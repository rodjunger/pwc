package pwc

import (
	"testing"
)

func TestCapitalizeFirstLetter(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"world", "World"},
		{"123", "123"},
		{"", ""},
		{"ab cd ef", "Ab Cd Ef"},
	}

	for _, test := range tests {
		result := capitalizeFirstLetter(test.input)
		if result != test.expected {
			t.Errorf("Expected CapitalizeFirstLetter(%q) to be %q, but got %q", test.input, test.expected, result)
		}
	}
}

func TestCapitalizePhrases(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello. how are you? i'm fine, thank you.", "Hello. How are you? I'm fine, thank you."},
		{"hello. world!", "Hello. World!"},
		{"hello", "Hello"},
		{"312. 111?", "312. 111?"},
		{"", ""},
	}

	for _, test := range tests {
		result := CapitalizePhrases(test.input)
		if result != test.expected {
			t.Errorf("Expected CapitalizePhrases(%q) to be %q, but got %q", test.input, test.expected, result)
		}
	}
}
