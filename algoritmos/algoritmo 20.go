package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Print("Digite uma string em camelCase: ")
	fmt.Scanln(&str)

	isCamelCase := true
	wordCount := 1

	// Verifica se a primeira letra é maiúscula
	if len(str) > 0 && str[0] >= 'a' && str[0] <= 'z' {
		isCamelCase = false
	}

	// Percorre a string a partir do segundo caractere
	for i := 1; i < len(str); i++ {
		// Verifica se o caractere é maiúsculo
		if str[i] >= 'A' && str[i] <= 'Z' {
			// Se o caractere anterior não for uma letra maiúscula, significa que uma nova palavra começou
			if str[i-1] < 'A' || str[i-1] > 'Z' && str[i-1] < 'a' || str[i-1] > 'z' {
				isCamelCase = false
			}
			wordCount++
		}
	}

	if isCamelCase {
		fmt.Println("A string está em camelCase.")
		fmt.Printf("A quantidade de palavras na string é: %d\n", wordCount)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}
