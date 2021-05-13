package main

import "fmt"

func main() {
	/*
		Initial entry point, sets arrays and slices, and then finds the difference between the last 2 items in the slices
		Return: None
	*/

	// Declaring array in which slices will be made from
	arr := [6]int{3444, 2342, 5234, 234, 412, 764}

	// Getting length of arr array
	arr_length := len(arr)
	arr_middle := arr_length / 2

	// Creating new slices with new variables
	arr_slice := arr[0:arr_middle]
	arr_slice1 := arr[arr_middle:arr_length]

	// Printing parent array
	fmt.Printf("\nParent array: ")
	fmt.Println(arr, "\n")

	// Printing slices
	fmt.Println("Slices:")
	fmt.Println(arr_slice)	
	fmt.Println(arr_slice1, "\n")	

	// Printing last value
	fmt.Println("Last value:")
	fmt.Println(arr_slice[len(arr_slice)-1])	
	fmt.Println(arr_slice1[len(arr_slice1)-1], "\n")

	// Printing difference of the last items in slices
	fmt.Printf("The difference of the integers above: %d", arr_slice[len(arr_slice)-1] - arr_slice1[len(arr_slice1)-1])
}