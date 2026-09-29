package main

import "fmt"

// * -> puntero
// & -> dirección de memoria
func incrementar(numero *int) {
	*numero++
}
func punteros() {
	valor := 10
	fmt.Println("Valor original:", valor)

	incrementar(&valor)

	fmt.Println("Valor después de incrementar (con puntero):", valor)

	puntero := new(int)

	fmt.Println("Valor del puntero:", *puntero)

	*puntero = 20

	fmt.Println("Valor del puntero después de asignar 20:", *puntero)
}
