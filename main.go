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

	// Working with collections
	// Arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[2])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//Slices
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4, 5, 6, 7)
	fmt.Println("slice is", slice)

	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]

	fmt.Println(s2, s3, s4)

	// Maps
	m := map[string]int{
		"foo": 42,
		"bar": 69,
	}
	fmt.Println(m)
	delete(m, "bar")
	fmt.Println(m)

	//Structs
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 33
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)

	v := user{
		ID:        69,
		FirstName: "Zaphod",
		LastName:  "Beeblebrox",
	}
	fmt.Println(v)
}
