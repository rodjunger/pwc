package main

import (
	"fmt"
	"os"

	challanges "github.com/rodjunger/pwc/challanges"
)

var functions = map[string]func(string) string{
	"1": challanges.ReverseWordOrder,
	"2": challanges.RemoveDuplicateLetters,
	"3": challanges.LongestPalindromeSubstring,
	"4": challanges.CapitalizePhrases,
	"5": challanges.IsAnagramOfPalindrome,
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./pwc [challange number] [challange input]")
		return
	}
	challangeNumber, challangeInput := os.Args[1], os.Args[2]

	if _, ok := functions[challangeNumber]; !ok {
		fmt.Println("Invalid challange ID")
		return
	}

	fmt.Println(functions[challangeNumber](challangeInput))
}
