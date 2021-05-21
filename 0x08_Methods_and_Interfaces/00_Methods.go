package main

import "fmt"

// Creating struct for car
type car struct {
	name string
	horsepower int
	rating float64
}

// Creating method "print_name" for car
func (the_car car) print_name() {
	fmt.Printf("This car models name is: %v\n", the_car.name)
}

func main() {
	/*
		Initial entry point, creates and utilizes a method for car
		Return: None
	*/

	// Creating new car instances
	new_car := car{"Mustang", 310, 7.8}
	Another_one := car{"Charger", 425, 7.4}
	// All of the available methods and values for car
	// car.name
	// car.horsepower
	// car.rating
	// car.print_name()

	// Using the cars instance method print_name to print cars name
	new_car.print_name()
	Another_one.print_name()
}
