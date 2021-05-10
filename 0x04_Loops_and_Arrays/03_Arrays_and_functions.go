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

	// Declaring arrays without setting size, the same will be accepted by the function
	arr3 := []int{32, 64, 234, 324, 235}
	arr4 := []int{3245, 45, 467, 9875, 123, 564, 1235, 135}

	// Setting array choices, these will be used for both cases
	arr1_choice := 3
	arr2_choice := 5
	
	// Printing results of array functions
	fmt.Printf("The sum of arr1[%d] and arr2[%d] is: %d\n", arr1_choice, arr2_choice, Array_Math(arr1, arr2, arr1_choice, arr2_choice))
	fmt.Println("")
	fmt.Printf("The sum of arr3[%d] and arr4[%d] is: %d\n", arr1_choice, arr2_choice, Array_Math_noSize(arr3, arr4, arr1_choice, arr2_choice))
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

func Array_Math_noSize(arr3 []int, arr4 []int, a int, b int) int{
	/*
		Adds 2 random index values together from int arrays passed in, however the arrays or not size specific
		Return: int	
	*/

	// Printing arrays passed in to the function
	fmt.Println(arr3)
	fmt.Println(arr4)

	// Returning sum of 2 array index's
	return arr3[a] + arr4[b]
}
