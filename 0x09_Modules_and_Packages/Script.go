package main

import (
	"fmt"
	"github.com/Syssos/hello/example"
)

// Declaring constant variable
const Greeting = "Hello Go writers!"

// Initializing main package
func init() {
	/*
		Initial entry point, Prints greeting
		Return: None
	*/

	// Printing greeting
	fmt.Println(Greeting)
}

func main() {
	/*
		Programs starting point, calls example package to print to screen
		Return: None
	*/

	// Printing text to screen via package
	example.Example_func()
}