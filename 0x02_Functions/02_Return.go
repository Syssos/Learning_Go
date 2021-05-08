package main

import (
	"math"
	"fmt"
)

func main(){

	// declaring a float, because abs requires that type
	number := float64(-2)

	// Here we are using our function in the process of creating a new value, not this is also where the abs value is printed
	abs_num := math_abs(number) * number

	// This will print the new number we created with our function
	fmt.Printf("The value of \"number\" multiplied by its absolute value is: %g\n", abs_num)
}

func math_abs(x float64) float64 {

	// finds abs of value passed in and assigns it to the variale
	abs_val := math.Abs(x)

	// Prints the abs value calculated above.
	fmt.Printf("The abs of \"number\" is: %g\n", abs_val)

	// Returns the abs value to be later used in our code
	return(abs_val)
}