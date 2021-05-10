package main

import "fmt"

func main() {
	/*
		Initial entry point, prints greeting n amount of times
		Return: None
	*/

	// Declaring greeting variable to be printed
	greeting := "Hello World!"
	n := 10

	// For Loop starting with i = 0
	fmt.Println("Results for: i = 0")
	for i := 0; i < n; i++ {
		// example output: Hello World!, Greeting number: 0/10
		fmt.Printf("%v, Greeting number: %d/%d\n", greeting, i, n)
	}

	// Printing newline for spacing
	fmt.Println("")

	// For loop starting with i = 1
	fmt.Println("Results for: i = 1, with n + 1 in condition")
	for x := 1; x < n + 1; x++ {
		// example output: Hello World!, Greeting number: 0/10
		fmt.Printf("%v, Greeting number: %d/%d\n", greeting, x, n)
	}
}