package main

import "fmt"

func main() {
	/*
		Initial entry point, demenstrates removing an item from a slice.
		Return: None
	*/

	// Declaring parent array
	arr := []int{312, 1243, 123, 654, 345, 865}

	// Printing parent array
	fmt.Println(arr)

	// Removing arr[1], note this will not keep the array in order
	// Replacing value with last item in array
	arr[1] = arr[len(arr)-1]

	// clearing the value held in last array item
	arr[len(arr)-1] = 0

	// Truncating slice
	arr = arr[:len(arr)-1]

	// Printing updated arr
	fmt.Println(arr)
}