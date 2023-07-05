package challanges

// IsAnagramOfPalindrome retorna "true" se a string in pode ser rescrita como um palindromo, e "false" caso contrário
// A ideia é simplesmente contar o numero de caracteres com quantidade impar na string, se tivermos mais de um,
// Então é impossivel criar um palindromo, a contagem de pares não importa porque sempre podemos encaixar
// quantidades iguais na "direita" e na "esquerda" do palindromo para casar os caracteres
func IsAnagramOfPalindrome(in string) string {
	countOfRunes := map[rune]int{}
	for _, character := range in {
		countOfRunes[character] += 1
	}

	cnt := 0

	for _, v := range countOfRunes {
		if v%2 != 0 {
			cnt += 1
			if cnt > 1 {
				return "false"
			}
		}
	}
	return "true"
}
