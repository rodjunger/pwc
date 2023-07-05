package challanges

import "testing"

func TestIsAnagramOfPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"racecar", "true"},
		{"rraacec", "true"},
		{"", "true"},
		{"a", "true"},
		{"abc", "false"},
		{"aab", "true"},
		{"aabb", "true"},
		{"aabcc", "true"},
		{"aabccd", "false"},
		{"aaabbcc", "true"},
		{"aaabbccd", "false"},
		{"tactcoa", "true"},
		{"tactcoaa", "false"},
	}

	for _, test := range tests {
		result := IsAnagramOfPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsAnagramOfPalindrome(%q) returned %q, expected %q", test.input, result, test.expected)
		}
	}
}
