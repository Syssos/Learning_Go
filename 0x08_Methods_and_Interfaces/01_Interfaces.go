package main

import (
	"fmt"
	"math"
)
// Using math to create global Pi variable so I dont have to use math.Pi in code
const Pi = math.Pi

// Declaring struct for square
type sqr struct {
	width, height float64
}

// Creating methods for square
func (s sqr) area() float64{
	// Returning sum of width and height
	return s.width * s.height
}

func (s sqr) per() float64{
	// Returning perimeter of square
	return 2 * s.width + 2 * s.height
}

// Declaring struct for Circle
type crcl struct {
	radius float64
}

// Creating methods for circle
func (c crcl) area() float64 {
	// Return same value as Pir^2
	return Pi * c.radius * c.radius
}

func (c crcl) per() float64 {
	// Returning circumference or perimeter of a circle
	return 2 * Pi * c.radius
}

// Creating interface
type geometry interface {
    area() float64
    per() float64
}

func main(){
	/*
		Initial entry point, Creates 2 structs and utilizes interface
		Return: None
	*/

	// Creating sqr struct
	r := sqr{width: 3, height: 4}
	// Available methods and values
	// sqr.width
	// sqr.height
	// sqr.area()
	// sqr.per()

	// Creating crcl struct
	c := crcl{radius: 3}
	// Available methods and values
	// crcl.radius
	// crcl.area()
	// crcl.per()
	
	// Printing structs above
	fmt.Println("sqr:", r)
	fmt.Println("crcl:", c, "\n")

	// geometry interface creation w/ crcl struct instance
	g := geometry(c)
	// Available methods
	// g.area()
	// g.per()
	// g is essentually instance of either crcl or sqr

	// Printing Geometry interface
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.per(), "\n")

	// Passing r/c structs to function, resulting in geometry instance being created and used
	doMath(r)
	fmt.Println()
	doMath(c)
}

func doMath(g geometry) {
	/*
		Uses geometry interface to print area and perimeter of shape
		Return: None
	*/
	fmt.Printf("Area: %.02f\n", g.area())
	fmt.Printf("Perimeter: %.02f\n", g.per())
}