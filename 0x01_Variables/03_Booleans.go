package main

import "fmt"

func main() {

	// Setting are bool variables
	var booliscool bool = true
	var boolisntcool bool = false


	if booliscool {

		fmt.Printf("Booliscool: %t\n", booliscool)
	}

	if boolisntcool {

		// We shouldn't see this in our output
		fmt.Printf("First boolisntcool: %t\n", boolisntcool)
	}

	boolisntcool = true

	if boolisntcool {

		// Now that the variable is set to true this statement should print
		fmt.Printf("Second boolisntcool: %t\n", boolisntcool)
	}
}