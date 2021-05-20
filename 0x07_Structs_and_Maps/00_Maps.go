package main

import "fmt"

func main() {

    // Creating map varaibles, both with make and composite litteral
    m := make(map[string]int)
    m2 := map[string]bool{}

    // Setting key, value pairs for maps
    m["k1"] = 7
    m["k2"] = 13
    m2["Awesome"] = true
    m2["Lame"] = false
    m2["Maps"] = true

    // Printing Initial maps
    fmt.Println("Map created with make:", m)
    fmt.Println("Map created with composite literal:", m2)

    // showing length of both maps
    fmt.Println("m len:", len(m))
    fmt.Println("m2 len:", len(m2))

    // Delete values from maps
    delete(m, "k2")
    delete(m2, "Lame")

    // Showing values were deleted
    fmt.Println("updated m map:", m)
    fmt.Println("updated m2 map:", m2)
}