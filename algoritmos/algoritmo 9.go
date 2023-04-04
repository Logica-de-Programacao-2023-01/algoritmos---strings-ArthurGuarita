package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, retire, add string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)

	fmt.Print("Digite o caractere a ser substituido: ")
	fmt.Scan(&retire)

	fmt.Print("Digite o caractere a ser add: ")
	fmt.Scan(&add)
	//troca
	fmt.Print(strings.ReplaceAll(str, retire, add))

}
