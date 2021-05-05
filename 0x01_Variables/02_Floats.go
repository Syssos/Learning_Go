package main

import "fmt"

func main() {

	// Declaring the float variables, holding the value of pi.
	var pi float32 = 3.14
    var lng_pi float64 = 3.1415926535

    fmt.Printf("Pi: %g\nLong Pi: %g\n", pi, lng_pi)

    // We have the short version of pi, within the long version, lets use lng_pi to print the value we have for pi
    fmt.Printf("Pi, using lng_pi: %.2f\n", lng_pi)
    fmt.Println("")

    // Like in 01_Ints, lets see how the dataypes are stored
    fmt.Printf("Pi: %T\nLong Pi: %T\n", pi, lng_pi)
}