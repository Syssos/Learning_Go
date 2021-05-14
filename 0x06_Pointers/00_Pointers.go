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

	// Printing the derefrenced 
	fmt.Println(*le_point)

	// Printing the Pointers address, derefrenced variables address, and the variables address
	fmt.Printf("Pointer address:%p\nDerefrenced pointers address: %p\nVariables address: %p\n", &le_point, *&le_point, &str)
}