package main

import (
	"fmt"
)

func suma(a int, b int) int {
	return a + b
}

func dividirV1(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero no permitida")
	}
	return a / b, nil
}

// Funcion que devuelve mas de 1 valor
func dividirV2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero no permitida")
	}
	return a / b, nil
}

// función con número variable de argumentos
func imprimirNombres(nombres ...string) {
	for _, nombre := range nombres {
		fmt.Println("Nombre:", nombre)
	}
}

// Clausure
func contador() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Estructuras
type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func parte3() {
	cociente, error := dividirV2(4, 2)
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Cociente:", cociente)
	}
	cont := contador()
	fmt.Println("Contador:", cont())
	fmt.Println("Contador:", cont())
	fmt.Println("Contador:", cont())
	fmt.Println("Contador:", cont())

	rect := Rectangulo{Ancho: 5, Alto: 10}
	fmt.Println("Área del rectángulo:", rect.Area())
}
