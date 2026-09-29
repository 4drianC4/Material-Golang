package main

import "fmt"

func parte2() {
	// Defer
	defer fmt.Println("Esto se ejecuta al final de la función")
	// Condicionales
	edad := 20
	if edad < 18 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("Eres mayor de edad")
	}

	// Switch
	dia := "Lunes"
	switch dia {
	case "Lunes":
		fmt.Println("Hoy es lunes")
	case "Martes":
		fmt.Println("Hoy es martes")
	default:
		fmt.Println("Hoy es otro día")
	}

	// Bucles
	for i := 0; i < 5; i++ {
		fmt.Printf("Número: %d\n", i)
	}

	n := 0
	for n < 5 {
		fmt.Printf("Número en bucle while: %d\n", n)
		n++
	}

	// Range en slices
	slice := []string{"a", "b", "c"}
	for index, value := range slice {
		fmt.Printf("Índice: %d, Valor: %s\n", index, value)
	}
}
