package pwc

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aba", true},
		{"abba", true},
		{"abc", false},
		{"saippuakivikauppias", true},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("isPalindrome(%q) returned %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestLongestPalindromeSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"babad", "bab"},
		{"", ""},
		{"abcdefghijkl", "a"},
		{"jfoaijdfoashjf9281hfcivicj1dlksja98dj1", "civic"},
		{"saippuakivikauppias", "saippuakivikauppias"},
	}

	for _, test := range tests {
		result := LongestPalindromeSubstring(test.input)
		if result != test.expected {
			t.Errorf("LongestPalindromeSubstring(%q) returned %q, expected %q", test.input, result, test.expected)
		}
	}
}
