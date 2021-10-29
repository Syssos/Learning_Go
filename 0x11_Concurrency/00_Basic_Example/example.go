package main

import (
	"fmt"
	"time"
)

func delay_print(str string){
	time.Sleep(8 * time.Second)
	fmt.Println(str)
}

func main() {
	fmt.Println("Starting")

	go delay_print("From Go Routine")

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Not in Routine")
	}
}