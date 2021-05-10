package main

import "fmt"

func main() {
	/*
		Initial entry point, using array to keep track of integer data
		Return: None
	*/

	// Decaring array variable, 
	arr := [5]int{1, 0, 2342, 5763, 9074}
	name_arr := [5]string{"Bob", "Billy", "Sam", "Steven", "Scott"}

	// Printing int array, and name array
	fmt.Printf("\nInt array: ")
	fmt.Println(arr)
	fmt.Printf("\nName array: ")
	fmt.Println(name_arr)

	// Printing item in int array at index
	fmt.Printf("\nValue at index 2 in int array: %d\n", arr[2])

	// Printing sum of 2 values at different index's in int array
	fmt.Printf("\nThe sum of index 0 and 2 in int array: %d\n", arr[0] + arr[2])

	// Printing Greeting for each name
	fmt.Printf("\nGreeting for each name:\n")

	// Loop uses counter set the same as array length
	for i:=0; i<5; i++ {
		fmt.Printf("Hello %v\n", name_arr[i])
	}
}