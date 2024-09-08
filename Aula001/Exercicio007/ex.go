package main

import (
	"fmt"
)

func main() {

	valor1 := true
	valor2 := false
	valor3 := true

	fmt.Printf("&&: %v \n", valor1 && valor2 && valor3)
	fmt.Printf("||: %v \n", valor1 || valor2 || valor3)
	fmt.Printf("!valor1: %v \n", !valor1)
	fmt.Printf("!valor2: %v \n", !valor2)
	fmt.Printf("!valor3: %v \n", !valor3)
}
