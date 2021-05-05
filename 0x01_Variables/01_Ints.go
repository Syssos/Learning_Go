package main

import "fmt"

func main() {

	// Declaring variables, Each one has a datatype that close, but not the same.
	var c int8 = 127
	var o int16 = 1568
	var d int32 = 654351
	var e int64 = 35138484384

	var h uint8 = 255
	var a uint16 = 2563
	var ck uint32 = 532653
	var s uint64 = 313843543

	// Printing the actual value of each variable
	fmt.Printf("c: %d,\no: %d,\nd: %d,\ne: %d,\nh: %d,\na: %d,\nck: %d,\ns: %d\n", c, o, d, e, h, a, ck, s)
	fmt.Println("")

	// Printing the dataype of each variable
	fmt.Printf("c: %T,\no: %T,\nd: %T,\ne: %T,\nh: %T,\na: %T,\nck: %T,\ns: %T\n", c, o, d, e, h, a, ck, s)
}