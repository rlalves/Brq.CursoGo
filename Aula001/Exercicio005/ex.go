package main

import (
	"fmt"
)

func main() {

	valor := 10

	valor += 10
	fmt.Printf("soma: %d \n", valor)

	valor -= 10
	fmt.Printf("subtracao: %d \n", valor)

	valor *= 10
	fmt.Printf("multiplicacao: %d \n", valor)

	valor /= 10
	fmt.Printf("divisao: %d \n", valor)

	valor %= 10
	fmt.Printf("resto: %d \n", valor)
}
