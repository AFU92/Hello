package main

import (
	"fmt"
	"time"
)

// Comentario sencillo xD

func main() {
	// var -> variable
	// int -> tipo de dato intero
	// = -> asignación

	var sum int = 8 + 9
	var rest int = 8 - 9
	var name string = "Andrea"
	var last_name string = "Fuentes"

	// Actualizar la variable
	name = "Mary"

	country := "Colombia"

	age := 27

	// Concatenar strings
	fmt.Println(".....Hello " + name + " " + last_name + ".....\n")

	fmt.Println("...Country:\n" + country + "\n\n...Age: ")
	fmt.Println(age)

	// Tiempo de espera un segundo
	time.Sleep(time.Second * 1)

	fmt.Println("\n8+9 =")
	fmt.Println(sum)

	// Tiempo de espera dos segundos
	time.Sleep(time.Second * 2)
	fmt.Println("\n8-9 =")
	fmt.Println(rest)
	/*
		Comentario de
		varias lineas */

}
