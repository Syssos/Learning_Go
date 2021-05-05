package main

import "fmt"

func main() {

	// Declaring the variable x is a string, and then assigning the value "Hello" to it
	var x string = "Hello"
	// Declaring another string variable, with a different value.
	var y string = "World!"

	fmt.Println("Strings x and y:")
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n" ,y)

	fmt.Println("")

	// Creating a new string variable out of 2 other string variables
	var z string = x + y

	fmt.Println("New string created from x and y:")
	fmt.Println(z)

	fmt.Println("\nWait we dont have a space\n")
	// Wait we need a space in between Hello and World

	// Adding and string containing a space in the middle of the two variables acheives this
	var z_with_space string = x + " " + y

	// Now we should see "Hello World!"
	fmt.Println("Thats much better")
	fmt.Println(z_with_space)
}