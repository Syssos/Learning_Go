package main

import "fmt"

func main() {

	Firstname := "James"
	Lastname := "Dean"

	Greet_user(&Firstname, &Lastname)

	fmt.Printf("Goodbye, %v %v\n", Firstname, Lastname)

	fmt.Println("This should've been")
	fmt.Println("Goodbye, James Dean")
}

func Greet_user(Firstname *string, Lastname *string) {

	*Firstname = "Timmy"
	*Lastname = "Turner"

	fmt.Printf("Hello %v %v, How are you?\n", *Firstname, *Lastname)
}