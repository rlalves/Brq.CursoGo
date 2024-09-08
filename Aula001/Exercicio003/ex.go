package main

import (
	"fmt"
)

func main() {

	variavelInteiro := 50
	var variavelFloat float64
	var variavelString string

	variavelFloat = float64(variavelInteiro)
	variavelString = fmt.Sprint(variavelInteiro)

	fmt.Printf("Variavel int: %d \n", variavelInteiro)
	fmt.Printf("Variavel float: %f \n", variavelFloat)
	fmt.Printf("Variavel string: %s \n", variavelString)
}
