package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str1)
	fmt.Print("Digite outra string: ")
	fmt.Scan(&str2)
	if strings.Contains(str1, str2) {
		fmt.Print("É um anagrama. ")
	} else {
		fmt.Print("Não são anagramas. ")
	}
}
