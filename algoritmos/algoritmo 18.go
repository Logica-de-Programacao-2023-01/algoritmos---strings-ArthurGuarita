package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	isNumber := true
	for _, char := range str {
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
