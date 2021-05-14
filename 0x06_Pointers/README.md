# Pointers

Pointers are an amazing little method to save the address of a variable. This can be helpful for multiple reasons, but to understand lets go back to how variables are passed in functions.

When working with a function we declare any arguemts and the arguments type in the function decloration. This is because when passing a function a variable, the variable data is used, to create a new instance.

## 00_Pointers

To better understand pointers lets break down this example using more code.

### code
``` go
package main
import "fmt"
func main (){
	str := "This is a test"
	fmt.Printf("<%p> %v\n", &str, str)

	no_point(str)
}

func no_point(new_str string) {
	fmt.Printf("<%p> %v", &new_str, new_str)
}
```
### output

**<0xc000010240> This is a test
<0xc000010260> This is a test**

If the code above is ran we will see 2 different memory address being printed, aside the same string. This is because the function no_point created a new variable, new_str, and stored the data it was passed into it.

Some may ask why isn't the variable, and its location passed to the function? The reason for that is to prevent data from accidentally being overwritten. If a copy is made and returned then there is less of a chance of the original data source to be modified when unwanted.

Here is a way we can utilize pointers to have the function use the same variable we created in the main function

### code
```go
package main
import "fmt"
func main() {
	str := "This is a test"
	fmt.Printf("<%p> %v", &str, str)

	wth_point(*str)
}

func wth_point(*new_str string) {
	// the value of new_str is the address of strT
	// To use str in this function we need to derefrence it with * 
	fmt.Printf("<%p> %v", *&new_str, *str)
}
```

### output

**<0xc00010a050> This is a test
<0xc00010a050> This is a test**

The code above is starting to look silmular to what we saw in the 00_Pointers example, however we are pasing the variable through a function.

Pointers are not something you will always use when writing Go code, however they are important to understand as there will be times they are needed to acheive a softwares end goal.

## Conclusion

While the use of a pointer can be pretty straight forward, knwoing when to utilize them can make all the difference. Using pointers helps keep the program light weight as it can reduce the amount of memory used by the program, however this introduces problems that can be tricky to track down if one is not paying attention.