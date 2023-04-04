package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)

	fmt.Println(strings.ToUpper(str))
}
