package main

import (
	"fmt"
)

//
type Cap struct {
	// Modelos de datos
	brand string
	color string
	price float32
	curve bool
}

func main() {

	var cap_black = Cap{
		brand: "Nike",
		color: "Black",
		price: 20.25,
		curve: true}

	var cap_red = Cap{
		"Adidas",
		"Black",
		24.50,
		true}

	fmt.Println(".....Black cap")
	fmt.Println(cap_black.brand)
	fmt.Println(cap_black.price)

	fmt.Println(".....Red cap")

	fmt.Println(cap_red)

	var num1 float32 = 10
	var num2 float32 = 6

	var num3 int = 10
	var num4 int = 6

	fmt.Print("\nThe sum is: ")
	fmt.Println(num1 + num2)

	fmt.Print("The rest is: ")
	fmt.Println(num1 - num2)

	fmt.Print("The mult is: ")
	fmt.Println(num1 * num2)

	fmt.Print("The div float is: ")
	fmt.Println(num1 / num2)

	fmt.Print("The div int is: ")
	fmt.Println(num3 / num4)

	fmt.Print("The residue is: ")
	fmt.Println(num3 % num4)

}
