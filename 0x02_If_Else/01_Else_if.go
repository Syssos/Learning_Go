package main

import "fmt"

func main() {
	// Declaring Emotion varibale
	Emotion := "Happy"

	// If statements checks users mood. The else if allows us to have more emotion reactions.
	if Emotion == "Happy" {
		fmt.Println("Im happy with you!")
	} else if Emotion == "Sad" {
		fmt.Println("Don't be sad, were learning Go together!")
	} else if Emotion == "Angry" {
		fmt.Println("Don't be upet if you dont understand just yet, good things take time!")
	} else {
		fmt.Println("Emotion.error: Human emotion not recognized")
	}
}