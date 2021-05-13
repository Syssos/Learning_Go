package main

import "fmt"

func main() {
	/*
		Initial entry point, demenstrates removing an item from a slice.
		Return: None
	*/

	// Declaring parent slice
	slice := []int{312, 1243, 123, 654, 345, 865}

	// Printing parent slice
	fmt.Println(slice)

	// Removing slice[1], note this will not keep the slice in order
	// Replacing value with last item in slice
	slice[1] = slice[len(slice)-1]

	// clearing the value held in last slice item
	slice[len(slice)-1] = 0

	// Truncating slice
	slice = slice[:len(slice)-1]

	// Printing updated slice
	fmt.Println(slice)
}