package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, c string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)
	fmt.Print("Digite um caractere: ")
	fmt.Scan(&c)
	//
	fmt.Print(strings.Count(str, c))
}
