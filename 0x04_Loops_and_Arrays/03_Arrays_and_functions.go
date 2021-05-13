package main

import "fmt"

func main() {
	/*
		Initial entry point, prints results from math function, using arrays passed to it from main
		Return: None
	*/

	// Declaring arrays with set size, the function will need to accept arrays of the same size
	arr1 := [5]int{42, 3245, 32, 64, 234}
	arr2 := [6]int{81, 64, 234, 3245, 45, 467}

	// Setting array choices
	arr1_choice := 3
	arr2_choice := 5
	
	// Printing results of array functions
	fmt.Printf("The sum of arr1[%d] and arr2[%d] is: %d\n", arr1_choice, arr2_choice, Array_Math(arr1, arr2, arr1_choice, arr2_choice))
}

func Array_Math(arr1 [5]int, arr2 [6]int, x int, y int) int{
	/*
		Adds 2 random index values together from int arrays passed in
		Return: int	
	*/

	// Printing arrays passed in to the function
	fmt.Println(arr1)
	fmt.Println(arr2)

	// Returning sum of 2 array index's
	return arr1[x] + arr2[y]
}
