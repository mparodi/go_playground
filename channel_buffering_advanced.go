package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	nombre := make(chan string)
	edad := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println(<-edad)
		fmt.Println("Termine de procesar la edad")
		wg.Done()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println(<-nombre)
		fmt.Println("Termine de procesar el nombre")
		wg.Done()
	}()

	go func() {
		edad <- "18"
	}()

	nombre <- "buffered"

	wg.Wait()
}