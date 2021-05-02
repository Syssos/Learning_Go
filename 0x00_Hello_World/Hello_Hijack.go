package main

import "fmt"

func main() {

	Firstname := "Alexis"
	Lastname := "Texis"

	Greet_user(&Firstname, &Lastname)

	fmt.Printf("Goodbye, %v %v\n", Firstname, Lastname)

	fmt.Println("This should've been")
	fmt.Println("Goodbye, Alexis Texis")
}

func Greet_user(Firstname *string, Lastname *string) {

	*Firstname = "Riley"
	*Lastname = "Reid"

	fmt.Printf("Hello %v %v, How are you?\n", *Firstname, *Lastname)
}