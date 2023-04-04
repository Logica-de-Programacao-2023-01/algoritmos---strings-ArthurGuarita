package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//digitar as palavras

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite varias palavras: ")
	scanner.Scan()
	words := scanner.Text()

	//quantas palavras

	count := strings.Count(words, " ")
	fmt.Print(count + 1)
}
