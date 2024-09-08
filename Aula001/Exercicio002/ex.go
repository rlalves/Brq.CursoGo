package main

import "fmt"

func main() {

	variavelIntIniciado := 50
	var variavelIntNaoIniciado int

	variavelFloat64Iniciado := 50.5
	var variavelFloat64NaoIniciado float64

	variavelBoolIniciado := true
	var variavelBoolNaoIniciado bool

	variavelStringIniciado := "teste"
	var variavelStringNaoIniciado string

	fmt.Println("-----------------------------------------")
	fmt.Printf("Variavel int iniciado: %d \n", variavelIntIniciado)
	fmt.Printf("Variavel int não iniciado: %d \n", variavelIntNaoIniciado)
	fmt.Println("-----------------------------------------")

	fmt.Println("-----------------------------------------")
	fmt.Printf("Variavel float iniciado: %f \n", variavelFloat64Iniciado)
	fmt.Printf("Variavel float não iniciado: %f \n", variavelFloat64NaoIniciado)
	fmt.Println("-----------------------------------------")

	fmt.Println("-----------------------------------------")
	fmt.Printf("Variavel bool iniciado: %+v \n", variavelBoolIniciado)
	fmt.Printf("Variavel bool não iniciado: %+v \n", variavelBoolNaoIniciado)
	fmt.Println("-----------------------------------------")

	fmt.Println("-----------------------------------------")
	fmt.Printf("Variavel string iniciado: %s \n", variavelStringIniciado)
	fmt.Printf("Variavel string não iniciado: %s \n", variavelStringNaoIniciado)
	fmt.Println("-----------------------------------------")
}
