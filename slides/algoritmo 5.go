package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var n int
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)
	fmt.Printf("Digite um número: ")
	fmt.Scan(&n)
	var result string
	for i := 0; i < n; i++ {
		result = result + string(str[i])
	}
	fmt.Println(strings.ToUpper(result))
}
