package main

import "fmt"

func main() {

	Firstname := "James"
	Lastname := "Dean"

	Greet_user(Firstname, Lastname)

	fmt.Printf("Goodbye, %v %v\n", Firstname, Lastname)
}

func Greet_user(F string, L string) {

	F = "Timmy"
	L = "Turner"
	fmt.Printf("Hello %v %v, How are you?\n", F, L)
}