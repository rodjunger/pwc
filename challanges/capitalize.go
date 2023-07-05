package pwc

import (
	"unicode"
)

var phraseEndingLetters = map[rune]struct{}{
	'.': {},
	'!': {},
	'?': {},
}

func isPhraseEndingLetter(in rune) bool {
	_, ok := phraseEndingLetters[in]
	return ok
}

// Pula espaços e outros caracteres especiais antes de capitalizar
func capitalizeFirstLetter(in string) string {
	runes := []rune(in)
	for i, letter := range runes {
		if unicode.IsLetter(letter) {
			runes[i] = unicode.ToUpper(letter)
			break
		}
	}
	return string(runes)
}

// CapitalizePhrases coloca em maiúscula a primeira letra de cada frase na string in
func CapitalizePhrases(in string) string {
	var (
		capitalized   string = ""
		currentPhrase string = ""
	)

	for _, letter := range in {
		currentPhrase += string(letter)
		if isPhraseEndingLetter(letter) {
			capitalized += capitalizeFirstLetter(currentPhrase)
			currentPhrase = ""
		}
	}

	//Se sobrou alguma coisa, temos uma ultima frase que tem que ser capitalizada
	if currentPhrase != "" {
		capitalized += capitalizeFirstLetter(currentPhrase)
	}
	return capitalized
}
