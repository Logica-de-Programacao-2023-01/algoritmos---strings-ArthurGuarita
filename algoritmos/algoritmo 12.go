package main

import "fmt"

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)

	strInvert := ""
	for i := len(str) - 1; i >= 0; i-- {
		strInvert += string(str[i])
	}
	if strInvert == str {
		fmt.Printf("São palíndromos %s e %s", str, strInvert)
	} else {
		fmt.Print("Não são palíndromos")
	}
}
