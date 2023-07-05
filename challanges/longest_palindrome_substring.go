package pwc

func isPalindrome(in string) bool {
	stringLength := len(in)
	for i := 0; i < stringLength/2; i++ {
		if in[i] != in[stringLength-i-1] {
			return false
		}
	}
	return true
}

// LongestPalindromeSubstring acha a maior substring palindroma na string in, usando a função helper
// isPalindrome que verifica se uma string qualquer é palindroma
// Ideia: Sliding window na string indo do tamanho maior até o menor
func LongestPalindromeSubstring(in string) string {
	//Caso simples
	if in == "" {
		return ""
	}
	//i é o tamanho da substring que estamos vendo
	for i := len(in); i > 1; i-- {
		//Esse loop executa um numero de vezes correspondente a quantidade de substrings de tamanho i na string in
		for j := 0; j <= len(in)-i; j++ {
			if isPalindrome(in[j : i+j]) {
				return in[j : i+j]
			}
		}
	}
	return string(in[0])
}
