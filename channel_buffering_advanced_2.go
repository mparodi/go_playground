package main

import (
	"fmt"
	"sync"
)

type Auto struct{
	marca string
	modelo string
	persona Persona
}

type Persona struct{
	nombre string
	apellido string
}

func main() {
	// Cada vez que una goroutine espera por un channel se pausa su ejecucion

	autoChannel := make(chan Auto)
	personaChannel := make(chan Persona)
	var wg sync.WaitGroup

	wg.Add(2) //Agregamos 1 por cada entidad a computar

	go func() { // funcion de crear el auto
		var auto Auto
		auto.persona = <-personaChannel
		fmt.Println("Empiezo a crear el auto")
		auto.marca = "mercedes benz"
		auto.modelo = "amg"
		fmt.Println("Termine de crear el auto")
		autoChannel <- auto
		wg.Done()
	}()

	go func() {
		fmt.Println("Empiezo a crear la persona")
		var persona Persona
		persona.apellido = "Gomez"
		persona.nombre = "Felipe"
		fmt.Println("Termine de crear la persona")
		personaChannel <- persona
		wg.Done()
	}()

	var auto Auto
	auto = <- autoChannel
	wg.Wait()


	fmt.Println("Imprimo el auto creado")
	fmt.Println(auto.modelo)
	fmt.Println(auto.marca)
	fmt.Println(auto.persona.nombre)
	fmt.Println(auto.persona.apellido)
}