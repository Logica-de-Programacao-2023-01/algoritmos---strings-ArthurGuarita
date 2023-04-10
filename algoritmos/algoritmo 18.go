package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	isNumber := true
	for _, char := range input {
		if !unicode.IsDigit(char) {
			isNumber = false
			break
		}
	}

	if isNumber {
		fmt.Println("A string contém somente números.")
	} else {
		fmt.Println("A string não contém somente números.")
	}
}
