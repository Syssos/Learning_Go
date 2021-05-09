package main

import "fmt"

func main() {

	// Types can be automatically assigned to variables with ":="
	var greeting string = "Hello World"
	closing := "Goodbye"
	
	// If Go sets the wrong type for what we need the var for, we can change it via type conversion.
	x := -2
	float_x = float64(x)

	// If statements let us change the flow in the code, based on if the condition returns true or flase.
	if float_x > 0 {
		fmt.Println(greeting)
	} else if float x < 0 {
		fmt.Println(closing)
	} else {
		fmt.Println("How did I get here?")
	}
}