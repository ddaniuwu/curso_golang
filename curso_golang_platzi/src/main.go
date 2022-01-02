package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//CICLO FOR
func ciclofor() {
	//USANDO LIBRERIA TIME
	time.Sleep(time.Second * 5)
	//CICLO FOR
	fmt.Println("CICLO FOR")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

//CICLOFORWHILE
func cicloForWhile() {
	counter := 0
	for counter < 10 {
		fmt.Println("Desde counter:", counter)
		counter++
	}
}

func serverRequest() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hola mundo desde un server con GO")
	})

	server := http.ListenAndServe(":8080", nil)
	log.Fatal(server)

}

//COUNTER FOREVER
//counterForever := 0
//for{
//fmt.Println(counterForever)
//counterForever++
//}

func main() {
	//EJERCICIO DE OPERADORES ARITMETICOS
	fmt.Println("OPERACIONES")
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado", areaCuadrado)

	x := 10
	y := 30

	//Suma
	result := x + y
	fmt.Println("Suma", result)

	//Resta
	result = y - x
	fmt.Println("Resta", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	//Division
	result = y / x
	fmt.Println("Division", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo / residuo", result)

	ciclofor()
	cicloForWhile()

	// for i := 10; i > 0; i-- {
	// 	fmt.Println(i)
	// }

	serverRequest()
}
