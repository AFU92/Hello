package main

import (
	"fmt"
)

// Comentario sencillo xD

func main() {

	var num1 float32 = 10
	var num2 float32 = 6

	var num3 int = 10
	var num4 int = 6

	fmt.Print("The sum is: ")
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
