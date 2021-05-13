package main

import "fmt"

func main() {
	/*
		Initial entry point, prints 1 array and 2 slices for demenstration
		Return: None
	*/

	// Creating an array, in which the slice will be created from
	arr_array := [6]int{34, 35, 86, 12, 04, 21}

	// Creating 2 slices, from the same parent array
	arr_slice := arr_array[0:6] // Gives us entire array, as the index range will be 0-5
	arr_slice1 := arr_array[2:4] // Gives us 2nd and 3rd arr_array indexs

	// Printing array, and slicesd
	fmt.Println(arr_array)
	fmt.Println(arr_slice)
	fmt.Println(arr_slice1)
}