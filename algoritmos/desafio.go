package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, pattern string

	// Solicita ao usuário a string e o padrão
	fmt.Print("Digite a string: ")
	fmt.Scan(&str)
	fmt.Print("Digite o padrão: ")
	fmt.Scan(&pattern)

	// Utiliza a função strings,Index para encontrar os índices do padrão na string
	var indices []int
	start := 0
	for {
		index := strings.Index(str[start:], pattern)
		if index == -1 {
			break
		}
		indices = append(indices, start+index)
		start += index + len(pattern)
	}

	// Verifica se foram encontrados índices e exibe o resultado
	if len(indices) > 0 {
		fmt.Printf("O padrão '%s' ocorre nos seguintes índices da string: %v\n", pattern, indices)
	} else {
		fmt.Printf("O padrão '%s' não foi encontrado na string.\n", pattern)
	}
}
