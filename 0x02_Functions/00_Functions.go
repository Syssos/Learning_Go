package main

import "fmt"

func main() {

	// Creating variables, In this case a set of predetermanded saying
	greeting := "Hello there, how are you"
	task := "Today will be a busy day of learning go"
	depart := "Goodbye"

	// Greeting user
	greet_user(greeting)
	
	// Printing task lists 
	fmt.Println(task)

	// Closing statement for the user when leaving.
	user_depart(depart)
}

func greet_user(greeting string) {

	fmt.Printf("%v doing today?\n", greeting)
}

func user_depart(saying string) {

	fmt.Println(saying)
}