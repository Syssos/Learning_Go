package main

import "fmt"

func main() {
	// Setting emotion, if changed, the output of the script will change
	Emotion := "Happy"

	// Testing if the user is happy
	if Emotion == "Happy" {
		fmt.Println("Im glad your in a good mood!")
	} else { 
		fmt.Println("Please cheer up :(") // If the code got here it was because the Emotion was not Happy
	}
}