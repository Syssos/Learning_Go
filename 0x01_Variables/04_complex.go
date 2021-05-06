package main

import "fmt"

func main() {

	// Declaring variables, artimetics will be performed in this example and in order to do that they need to be of the same type
	var com1 complex64 = complex(10, 2)
	var com2 complex64 = complex(8, 4)
	var com3 complex128 = complex(5, 8)

	// Print data before arithmetics
	fmt.Println("Initial Complex variables:")
	fmt.Printf("com1: %g Type: %T\n", com1, com1)
	fmt.Printf("com2: %g Type: %T\n", com2, com2)
	fmt.Printf("com3: %g Type: %T\n", com3, com3)
	fmt.Println("")

	// Print Arithmetics
	fmt.Println("Complex basic Arithmetics")
	fmt.Printf("com1 + com2: %g\n", com1+com2)
	fmt.Printf("com1 - com2: %g\n", com1-com2)
	fmt.Printf("com1 / com2: %g\n", com1/com2)
	fmt.Printf("com1 * com2: %g\n", com1*com2)

}