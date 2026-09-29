package main

import (
	"fmt"
	"strings"
)

func main() {
	// Numeros
	entero := 10
	decimal := 3.14
	suma := entero + int(decimal)
	fmt.Println("Suma:", suma)

	// Texto
	mensaje := "Hola esto es un mensaje de texto en strings"
	concatenado := mensaje + " y se ha concatenado con este texto"
	fmt.Println("Texto concatenado:", concatenado)
	PasaraMayusculas := strings.ToUpper(mensaje)
	fmt.Println("Texto en mayúsculas:", PasaraMayusculas)

	// Booleanos
	verdadero := true
	falso := false
	fmt.Println("Valor verdadero:", verdadero)
	fmt.Println("Valor falso:", falso)
	negacion := !verdadero
	fmt.Println("Negación de verdadero:", negacion)

	// Arreglos
	arrayFijo := [2]int{1, 2}
	fmt.Println("Arreglo fijo:", arrayFijo)

	sliceVariable := []int{3, 4, 5}
	sliceVariable = append(sliceVariable, 6)
	fmt.Println("Slice variable después de agregar un elemento:", sliceVariable)

	// Longitud y capacidad
	fmt.Println("Longitud del slice:", len(sliceVariable))
	fmt.Println("Capacidad del slice:", cap(sliceVariable))

	// Mapas
	diccionario := map[string]int{
		"uno":    1,
		"dos":    2,
		"tres":   3,
		"cuatro": 4,
	}
	fmt.Println("Diccionario:", diccionario)

	// Estructuras
	type Persona struct {
		nombre string
		edad   int
	}
	persona := Persona{nombre: "Juan", edad: 30}
	fmt.Println("Estructura Persona:", persona)

	// Privada y Pública
	type Animal struct {
		especie string // Campo privado
		Nombre  string // Campo público
	}
	animal := Animal{especie: "Perro", Nombre: "Rex"}
	fmt.Println("Animal:", animal)
	// Acceso a campos públicos y privados
	fmt.Println("Nombre del animal (público):", animal.Nombre)
	// fmt.Println("Especie del animal (privado):", animal.especie
	// Esto causaría un error de compilación porque 'especie' es privado
	fmt.Println("Especie del animal (acceso indirecto):", animal.especie == "Perro")

	parte2()
	parte3()
}
