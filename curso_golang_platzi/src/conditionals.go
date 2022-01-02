package main

import "fmt"

func main() {
	//CONDICIONALES Y OPERADORES LOGICOS
	valor1 := 40
	valor2 := 30

	const miCadena = "String"

	if valor1 == valor2 {
		fmt.Println("Son iguales")
	} else {
		fmt.Println("No son iguales")
	}

	//OPERADOR &&
	if valor1 == 40 && miCadena == "String" {
		fmt.Println("Condicional aplicada con logica &&")
	}

	//ALGUNAS VECES EL CASE VA A ACOMPAÑADO DE LA CONDICION ANTES DE LA EVALUACION
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//SIN CONDICION
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	}
}
