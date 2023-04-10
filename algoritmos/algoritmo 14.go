package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Print("Digite uma string de números: ")
	fmt.Scan(&str)
	strconv.Atoi(str)

	var crescente bool
	crescente = true
	for i := 1; i < len(str); i++ {
		if str[i] > str[i-1] {
			crescente = false
			break
		}
	}
	if crescente {
		fmt.Print("Está em ordem decrescente.")
	} else {
		fmt.Print("Não está em ordem decrescente.")
	}
}
