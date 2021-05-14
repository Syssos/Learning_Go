package main

import "fmt"

func main() {
	/*
		Initial entry point, this program was designated to display how pointers work.
		Return: None

	*/

	// Declaring string variable to make pointer from
	str := "Howdy"

	// Declaring pointer by saving the memory address
	le_point := &str

	// Printing the variables value
	fmt.Printf("Variables value: %v\n", str)

	// Printing variables address
	fmt.Printf("Variables address: %p\n", &str)

	// Printing the value of pointer
	fmt.Printf("Pointers value: %v\n", le_point)

	// Printing the address of pointer
	fmt.Printf("Pointers address in memory: %p\n", &le_point)

	// Printing the derefrenced value
	fmt.Printf("Pointers derefrenced value: %v\n", *le_point)

	// Printing the type of pointer
	fmt.Printf("Pointers type: %T\n", le_point)
}