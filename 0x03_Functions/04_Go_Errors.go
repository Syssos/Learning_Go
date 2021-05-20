package main

import (
	"fmt"
	"errors"
)

func main() {
	/* 
		Initial entry point, Calls function to multiply w numbers, then checks for errs during function call
		Return: None
	*/

	// calling mulThese function expection error
	sum, err := mulThese(4, 0)
	if err != nil {
		// Print error if one occurs, err is not nil
		fmt.Println("Failed: ", err, " Sum Value: ", sum)
	} else {
		// Printing answer, err is nil
		fmt.Println("Answer: ", sum, " Err Value: ", err)
	}

	// calling mulThese function knowing no error occurs
	sum, err = mulThese(4, 4)
	if err != nil {
		// Printing error if err is not nil
		fmt.Println("Failed: ", err, " Sum Value: ", sum)
	} else {
		// Printing answer, err is nil
		fmt.Println("Answer: ", sum, " Err Value: ", err)
	}
}

func mulThese(x int, y int) (int, error) {
	/* 
		Multiplies 2 numbers together as long as either one of them is not 0
		Return: Sum of x and y, nil if no err
	*/

	// Checking that x and y are both not 0
	if x != 0 && y != 0 {
		return x * y, nil
	} else {
		// Returning error because either x or y was 0
		return -1, errors.New("Number passed cannot be 0")
	}
}