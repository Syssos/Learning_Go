package main

import "fmt"

func main(){
	/*
		Initial entry point, demestrates appending an item to a slice or array.
		Return: None
	*/

	// Declaring parent array
	arr := []int{234, 634, 345, 362, 435}

	// We want to end with a slice containing the first and last item from the parent array using and appending method
	// Starting slice with first variable
	arr_slice := []int{arr[0]}

	// Appending the last item in arr to the slice
	arr_slice = append(arr_slice, arr[len(arr)-1])

	// Printing parent array
	fmt.Printf("\nParent array: ")
	fmt.Println(arr, "\n")

	// Printing slice
	fmt.Printf("New slice: ")
	fmt.Println(arr_slice)

}