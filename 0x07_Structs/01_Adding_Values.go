package main

import "fmt"

type car struct {
	name string
	horsepower int
	rating float64
}

func main() {
	new_car := car{}

	// Printing empty new car
	fmt.Println(new_car)

	// Adding name & printing updated struct
	new_car.name = "Mustang"
	fmt.Println(new_car)

	// Creating variables horsepower will rely on
	engine := 230
	pro_pack := 50
	turbo := 30

	// Setting horsepower variable & printing
	new_car.horsepower = engine + pro_pack + turbo
	fmt.Println(new_car)

	// To keep things simple lets just pretend this variable is set like the comment below
	// new_car.rating = get_rating(new_car.name, new_car.horsepower)
	new_car.rating = 7.8

	// Printing "completed" new_car struct
	fmt.Println(new_car)
}