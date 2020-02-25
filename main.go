package main

import "fmt"

func main() {
	// Primitives
	var entero int
	entero = 42
	fmt.Println(entero)

	var float float32 = 3.14
	fmt.Println(float)

	Nombre := "Manolo"
	fmt.Println(Nombre)

	var b bool = false
	fmt.Println(b)

	// pointers
	var firstName *string = new(string)
	*firstName = "Manolo"
	fmt.Println(*firstName)

	lastName := "Lopez"
	fmt.Println(lastName)
	ptr := &lastName
	fmt.Println(ptr, *ptr)

	// Constants
	const pi = 3.1416
	fmt.Println(pi + 3)
	fmt.Println(pi + 1.2)

}
