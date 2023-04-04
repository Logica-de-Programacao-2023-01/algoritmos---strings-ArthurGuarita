package main

import "fmt"

func main() {
	var a string
	fmt.Print("Digite a string: ")
	fmt.Scan(&a)
	fmt.Println(len(a))
}
