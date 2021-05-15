package main

import "fmt"

// Declaring the car struct, every car will follow this format, or contain these variables.
type car struct {
    name string
    horsepower  int
    rating float64
}

func main() {

	// Declaring a new car
	new_car := car{"Mustang", 310, 7.8}

	// Printing the struct
	fmt.Printf("Car stuct: %v\n", new_car)

	// Printing the name of the car
	fmt.Printf("Car name: %v", new_car.name)
}