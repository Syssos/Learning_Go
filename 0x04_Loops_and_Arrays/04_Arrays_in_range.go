package main

import "fmt"

func main() {
	/*
		Initial entry point, prints a value while array in range
		Return: None
	*/

	// Declaring array variables
	arr := [7]int{4345, 1231, 5235, 2342, 23235, 235, 233}

	// For loop iteration to end of array
	for i, v := range arr {
		// Prints string at each iteration
		fmt.Printf("Value at index %d is %d\n", i, v)
	}
}