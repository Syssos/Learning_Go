package main

import "fmt"

func main() {

	Firstname := "Alexis"
	Lastname := "Texis"

	Greet_user(Firstname, Lastname)

	fmt.Printf("Goodbye, %v %v\n", Firstname, Lastname)
}

func Greet_user(F string, L string) {

	F = "Riley"
	L = "Reid"
	fmt.Printf("Hello %v %v, How are you?\n", F, L)
}