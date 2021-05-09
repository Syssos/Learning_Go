package main

import (
	"fmt"
	"math"
	"github.com/endeveit/guesslanguage"
)

func main() {
	/*
		Program entry point, sets variables, controls program flow
		Return: No return values
	*/

	// Creating variables that will be used throughout, Here is were we utilize the github function
	number := -2
	greeting := "Hello, this greeting is in"

	// Checking greetings language
	lang, err := guesslanguage.Guess(greeting)

	// Checking if err, exists, this will be reviewed in the next directory
	if err != nil {
		fmt.Println(err)
	}

	// Printing the results of package imported from github
	fmt.Printf("%v %v\n", greeting, lang)

	// Notice the type conversion, Abs requires a float64 type and number is type int
	fmt.Printf("The absolute value of \"number\" is: %g\n", math.Abs(float64(number)))

}