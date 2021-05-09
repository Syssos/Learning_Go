package main

import "fmt"

func main() {

	// Delcation Emotion variable
	Emotion := "Happy"

	// Switch case, simliar to if but more condesed, perfect for testing a variable status.
	switch Emotion {
		case "Excited", "Happy":
			fmt.Println("Good emotions for the win!")
		case "Sad", "Gloomy":
			fmt.Println("Even rainy days have rainbows, things get better!")
		case "Angry", "Irritated":
			fmt.Println("Anger leads to anguish, deep breaths padiwan")
		default:
			fmt.Println("What is love, baby don't hurt me")
	}
}