# Desafio de Código PwC

## Requisitos

Certifique-se de ter Go instalado em seu sistema antes de executar o programa.

## Executando Testes automáticos

```shell
go test ./...
```

## Como Compilar e Executar

1. Clone o repositório:

```shell
git clone github.com/rodjunger/pwc/challanges
```

2. Compile o programa:

```shell
go build
```

3. Execute o programa com os argumentos corretos:

```shell
./challanges [número_do_desafio] [entrada_do_desafio]
```

Certifique-se de substituir `[número_do_desafio]` pelo número correspondente ao desafio que deseja executar e `[entrada_do_desafio]` pela entrada apropriada para o desafio selecionado.

## Funções Disponíveis

O programa contém as seguintes funções, cada uma resolvendo um desafio específico:

1. **ReverseWordOrder**: Inverte a ordem das palavras em uma frase.

2. **RemoveDuplicateLetters**: Remove letras duplicadas de uma frase.

3. **LongestPalindromeSubstring**: Encontra a maior substring palindrômica em uma frase.

4. **CapitalizePhrases**: Capitaliza a primeira palavra de cada frase de uma string.

5. **IsAnagramOfPalindrome**: Verifica se uma palavra é um anagrama de um palíndromo.
