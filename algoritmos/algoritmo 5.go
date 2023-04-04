package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)
	//
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Print("sim")
	} else {
		fmt.Print("nao", i)
	}
}
