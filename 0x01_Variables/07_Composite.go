package main

import "fmt"

func main() {

	// Delcaring a boolean, float, int, complex, and string
	cool := true
	x := 3.14
	y := 4535
	z := 10 + 7i
	greeting := "Hello World"

	// Printing variables and types (weird spacing is for easier to read output)
	fmt.Printf("%v  %g    %d  %g     %v\n", cool, x, y, z, greeting)
	fmt.Printf("%T  %T %T   %T  %T", cool, x, y, z, greeting)
}