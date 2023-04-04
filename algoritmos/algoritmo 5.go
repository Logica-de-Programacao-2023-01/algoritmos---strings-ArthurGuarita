package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)
	//verificação
	if _, err := strconv.ParseFloat(str, 64); err == nil {
		fmt.Print("Sim")
	} else {
		fmt.Print("Não")
	}
}
