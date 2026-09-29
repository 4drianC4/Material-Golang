package main

import (
	"fmt"
	"math"
)

// Las interfaces no se implementan manualmente
// las interfaces se cumplen

type Forma interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func imprimirArea(f Forma) {
	fmt.Printf("El área de la forma es: %.2f\n", f.Area())
}

func interfaces() {
	c := Circulo{Radio: 5}
	imprimirArea(c)

	var cualquiera interface{} = "esto es cualquier cosa"
}
