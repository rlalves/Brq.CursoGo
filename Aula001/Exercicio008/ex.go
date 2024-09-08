package main

import (
	"fmt"
)

func main() {

	valor := 1024
	fmt.Printf("valor antes: %v \n", valor)

	var ponteiro *int = &valor
	*ponteiro = 0
	fmt.Printf("valor depois: %v \n", valor)
}
