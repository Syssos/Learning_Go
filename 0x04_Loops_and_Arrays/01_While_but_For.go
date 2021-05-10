package main

import "fmt"

func main() {
	/*
		Initial entry point, displays the closest way form of a while loop go offers.
		Return: None

	*/

	// Declaring counter variable
	count := 0

	// using for with only condition to mimic while
	for count < 10 {
		// Message being repeated each iteration
		fmt.Printf("Current loop Itteration: %d\n", count)
		// Incrementing counter
		count++
	}
}