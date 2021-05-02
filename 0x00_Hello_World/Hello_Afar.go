package main

import "fmt"

func main() {

	Firstname := "James"
	Lastname := "Dean"

	Greet_user(Firstname, Lastname)
}

func Greet_user(F string, L string) {

	fmt.Printf("Hello %v %v, How are you?", F, L)
}