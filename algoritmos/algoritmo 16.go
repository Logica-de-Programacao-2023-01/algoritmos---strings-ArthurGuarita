package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scan(&str1)

	fmt.Println("Digite a segunda str0ing: ")
	fmt.Scan(&str2)
	//CONJUNTOS - ESTÁ CONTIDO / CONTÊM
	if strings.Contains(str1, str2) {
		fmt.Println("A segunda é substring da primeira")
	} else if strings.Contains(str2, str1) {
		fmt.Println("A primeira é substring da segunda")
	} else {
		fmt.Println("Null")
	}
}
