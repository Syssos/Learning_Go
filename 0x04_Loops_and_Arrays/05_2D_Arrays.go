package main

import "fmt"

func main() {
	/*
		Initial entry point, creates and prints 2 Dimensional array
		Return: None
	*/

	// Declaring 2d integar array
	arr := [2][3]int{
		{43, 64, 73},
		{23, 75, 70},
	}

	// Printing array, and the value at row 1 column 1
	fmt.Println(arr, "\n\n arr[1][1]: ", arr[1][1])
}