package main

import "fmt"

type car struct {
	name string
	horsepower int
	rating float64
}

func main() {

	// Creating new car instance
	new_car := car{"Mustang", 310, 7.8}

	// Printing car instance
	fmt.Println(new_car)

	// "Removing" the value
	new_car.rating = 0

	// Printing updated car instance
	fmt.Println(new_car)
}