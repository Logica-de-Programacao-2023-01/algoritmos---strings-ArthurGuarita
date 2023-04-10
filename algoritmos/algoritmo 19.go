package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)
	//inverter
	strInverter := ""
	for i := len(str) - 1; i >= 0; i-- {
		strInverter += string(str[i])
	}
	fmt.Print("A inversão da string é", strInverter)
}
