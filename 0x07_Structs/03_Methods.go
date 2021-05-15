package main

import "fmt"

// Creating structure
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

	// Creating new car instances
	new_car := car{"Mustang", 310, 7.8}
	Another_one := car{"Charger", 425, 7.4}

	// Using the cars instance method print_name to print cars name
	new_car.print_name()
	Another_one.print_name()
}
