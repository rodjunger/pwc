package pwc

// RemoveDuplicateLetters remove todas as letras duplicadas de uma string
func RemoveDuplicateLetters(in string) string {
	var (
		out  = ""
		seen = map[rune]bool{}
	)
	for _, letter := range in {
		if seen[letter] {
			continue
		}
		out += string(letter)
		seen[letter] = true
	}
	return out
}
