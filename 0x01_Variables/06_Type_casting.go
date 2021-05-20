package main

import "fmt"

func main() {

	// Declaring variables
	var x float64 = 42
	var y int8 = int8(x)
	var z int64 = int64(y)

	// Printing the value and datatype of each variable
	fmt.Printf("x: %g Type: %T\ny: %d Type: %T\nz: %d Type: %T\n", x, x, y, y, z, z)
}