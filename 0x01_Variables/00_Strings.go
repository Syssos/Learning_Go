package main

import "fmt"

func main() {

	// Declaring the variable x as a string, and then assigning the value "Hello" to it
	var x string = "Hello"

	// Declaring another string variable, with a different value.
	var y string = "World!"

	// Printing strings to sceen, notice the use of Println and Printf
	fmt.Println("Strings x and y:")
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n" ,y)

	// Adding a new line better spacing, Println ends every string with a new line character for us
	fmt.Println("")

	// Creating a new string variable out of 2 other string variables
	var z string = x + y

	// Printing the new string made by adding string x and y together 
	fmt.Println("New string created from x and y:")
	fmt.Println(z)

	// Wait we need a space in between Hello and World for it to look right
	fmt.Println("\nWait we dont have a space between hello and world\n")

	// Adding and string containing a space in the middle of the two variables acheives this
	var z_with_space string = x + " " + y

	// Now we should see "Hello World!"
	fmt.Println(z_with_space)
}