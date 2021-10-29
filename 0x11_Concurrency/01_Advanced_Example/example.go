package main

import (
	"fmt"
    "log"
    "sync"
	"net/http"
	"encoding/json"
)

const PROGRAMMING_URL string = "https://v2.jokeapi.dev/joke/Programming?type=single"
const PUN_URL string = "https://v2.jokeapi.dev/joke/Pun?type=single"
const MISC_URL string = "https://v2.jokeapi.dev/joke/Miscellaneous?type=single"

// Structure to hold response data from api
type JokeResponse struct {
	Error bool `json: "error"`
	Category string `json: "category"`
	Joke string `json: "joke"`
}

// Function to fetch joke from API
func FetchJoke(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("An error occurred attempting to fetch joke")
	}
	defer resp.Body.Close()
	
	var data JokeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal("An error occurred in json decode")
	}

	if !data.Error {
		return data.Joke
	}

	return ""
}

// Non-Concurrent API calls
func FetchThreeJokes() []string {
	var jokes []string

	// This was displayed in 3 lines over 1 to easier see each API request
	jokes = append(jokes, FetchJoke(PROGRAMMING_URL))
	jokes = append(jokes, FetchJoke(PUN_URL))
	jokes = append(jokes, FetchJoke(MISC_URL))

	return jokes
}

// Concurrent API calls
func FetchThreeJokesConcurrent() []string {
	var jokes []string
	var wg sync.WaitGroup
    wg.Add(3)

	go func() {
		defer wg.Done()
		jokes = append(jokes, FetchJoke(PROGRAMMING_URL))
	}()
	go func() {
		defer wg.Done()
		jokes = append(jokes, FetchJoke(PUN_URL))
	}()
	go func() {
		defer wg.Done()
		jokes = append(jokes, FetchJoke(MISC_URL))
	}()

	wg.Wait()
	return jokes
}

// Testing functionality of both approaches
func main() {

	// Non-Concurrent
	jokes := FetchThreeJokes()
	for _, joke := range jokes {
		fmt.Println(joke)
	}

	// Non-Concurrent
	jokesC := FetchThreeJokesConcurrent()
	for _, j := range jokesC {
		fmt.Println(j)
	}
}