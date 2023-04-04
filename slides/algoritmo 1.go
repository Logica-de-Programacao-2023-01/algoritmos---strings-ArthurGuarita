package main

import "fmt"

func main() {
	var a, b string
	fmt.Print("Digite a primeira string: ")
	fmt.Scan(&a)
	fmt.Print("Digite a segunda string: ")
	fmt.Scan(&b)

	str3 := a + " " + b
	fmt.Println(str3)
}
