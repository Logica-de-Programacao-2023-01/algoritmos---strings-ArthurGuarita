package main

import (
	"fmt"
	"unicode"
)

func main() {
	//string
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)

	//verificação
	for i := 0; i < len(str); i++ {
		if unicode.IsDigit(rune(str[i])) {
			fmt.Print("A string contêm número")
			return
		}
	}
	fmt.Print("Não contêm números")
}
