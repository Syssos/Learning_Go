package main

import "fmt"

func main() {
	/*
		Initial entry point, Demenstrates working at the start of the list.
		Return: None
	*/

	// Declaring array in which slices will be made
	arr := [6]int{432, 234, 864, 967, 846, 836}

	// Creating 2 slices from the parent array
	arr_slice := arr[0:3]
	arr_slice1 :=arr[3:6]

	// Printing array
	fmt.Printf("\nOriginal array: ")
	fmt.Println(arr,"\n")

	// Printing Slices
	fmt.Println("Array slices:")
	fmt.Println(arr_slice)
	fmt.Println(arr_slice1, "\n")

	// Printing First value of each slice
	fmt.Println("First value from each array:")
	fmt.Println(arr_slice[0])
	fmt.Println(arr_slice1[0])

	// Printing sum of these 2 numbers
	fmt.Printf("\nSum of integers above: %d\n", arr_slice[0] + arr_slice1[0])
}