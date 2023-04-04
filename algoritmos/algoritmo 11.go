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
	vogais := "aeiouAEIOU"
	for _, c := range vogais {
		str = strings.ReplaceAll(str, string(c), "")
	}
	fmt.Print(str)
}
