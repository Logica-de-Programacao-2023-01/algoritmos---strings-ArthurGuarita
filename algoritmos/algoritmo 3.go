package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)
	//
	var retire, add string
	fmt.Print("Digite o caractere a ser retirado da string: ")
	fmt.Scan(&retire)
	fmt.Print("Digite o caractere a ser add: ")
	fmt.Scan(&add)
	//
	fmt.Println(strings.ReplaceAll(str, retire, add))
}
