package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)
	//invert
	strInvert := ""
	for i := len(str) - 1; i >= 0; i-- {
		strInvert += string(str[i])
	}
	fmt.Print("A inversão da string é ", strInvert)
}
