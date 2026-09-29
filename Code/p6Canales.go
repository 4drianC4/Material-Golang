package main

import (
	"fmt"
	"sync"
	"time"
)

func decirHola(canal chan<- string) {
	time.Sleep(1 * time.Second)
	canal <- "Hola"
}

func imprimirMensaje(canal <-chan string) {
	mensaje := <-canal
	println(mensaje)
}

func canales() {
	canal := make(chan string)

	go decirHola(canal)

	imprimirMensaje(canal)

	canal2 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			canal2 <- i
		}
		close(canal2)
	}()

	for numero := range canal2 {
		println("Número recibido del canal:", numero)
	}

	// Mutex
	var contador int
	var mu sync.Mutex

	//Writer
	go func() {
		for i := 0; i < 5; i++ {
			mu.Lock()
			contador++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Reader
	for i := 0; i < 3; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				mu.RLock()
				fmt.Println("Contador actual:", contador)
				mu.RUnlock()
				time.Sleep(200 * time.Millisecond)
			}
		}()
	}

	time.Sleep(2 * time.Second) // Esperar a que los goroutines terminen
}
