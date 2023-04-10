package main

import "fmt"

func main() {
	var str string
	fmt.Print("Digite uma string:")
	fmt.Scan(&str)

	// verificação de caracteres únicos
	var unique string
	for i := 0; i < len(str); i++ {
		c := str[i]
		count := 0
		for j := 0; j < len(str); j++ {
			if c == str[j] {
				count++
			}
		}
		if count == 1 {
			unique += string(c)
		}
	}

	// imprime os caracteres únicos
	fmt.Println("Caracteres únicos:", unique)
}
