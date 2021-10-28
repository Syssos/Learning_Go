package main

import "fmt"

func Sum(x int, y int) int{
	return x + y
}

func Dif(x int, y int) int {
	return x - y
}

func main() {
	fmt.Println(Sum(4,5))
	fmt.Println(Dif(4,5))
}