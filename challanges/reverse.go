package challanges

import "strings"

// ReverseWordOrder inverte a ordem das palavras em uma frase, sem trocar a ordem das letras das palavras
func ReverseWordOrder(in string) string {
	var (
		out   = ""
		parts = strings.Split(in, " ")
	)
	for i := len(parts) - 1; i >= 0; i-- {
		out += parts[i]
		if i != 0 {
			out += " "
		}
	}
	return out
}
