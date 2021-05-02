# 0x00_Hello_World
### First steps into Go. Below are descriptions of each one of the files, and why they are here. I would like to note, that the purpose of these files are for visual representation of ideas or topics that I find important. Some of the files are used as examples, and are not how I would write code, but representations of what I should ensure I either stay away from or keep note of, as to not have faulty or improper code when in a development environment. I have experiance in C and Python going into this so I imagine there will be some simularities to the languages I am already familiar with but Will use to following pages to outline the those simulatities and differences.

## Hello_World
This File was used to get me started. I used it to find out how to run, and compile the code. As well as outline essential components I will need to understand such as packages, and imports.

```go
package main
```

The line above refrence's the current package in which we are working, because this is not an external module I am creating I will be using the main package.

```go
import "fmt"
```

The import line is importing the format package, this package is what gives us access to ```Println()``` and ```Printf()```, which prints text to the screen.

```go
func main (){...}
```

When creating a Go script the code execution starts within the main function, because I need no other function all of the work will be done within this one for now. Above I have the opening and closing brackets "{...}",the code within the function running is represented by the three periods "..."

```go
fmt.Println("Hello World!")
```
The Println function is a function within fmt that prints a string with basic format options such as a new line at the end of the string. In later files the use of Printf will be present for verbs within a string. More on that a little further down.

The lines of code above combined and put in a file called Hello_World.go gives us our first Go script. This file can then be either interpreted through the go run command, or it can be compiled into an executable. If your on windows thats an exe, on mac and linux it is an elf executable.

## Hello_V
This file is simular to the Hello World file, however this one is demenstrating the use of a "verb" within a string. A "verb" is a temparary variable that is placed in the string, that is then replaced with the value of a given argument when the string is printed to the screen. This function can print many different data types as arguments. Datatypes have not been covered in detail yet, so this wont be explained much here, however this is a topic that will be touched on in later projects as this print function can be very useful in debugging.

```go
fmt.Printf("%v, it works!", var)
```

Above we see the printf function being passed a string, and a variable name. Within the string there is a special character combonation, "%v". This is simuilar to how C denoted a verb within a string. The Percent sign (%) tells printf that this is the location to put the variable, and "v" which tells print to print the data in the default format. If we were to replace the "v" with something like a "p" then we would get a different value like the location in memory in which that variable is saved.

## Hello_Afar
Now that we know how to print variables and strings to the screen. Lets take a look at how we can use this in a practical sense. In this file we take a look at how to print stings, and variables in a function outside of main.

The main function in this file is used for 2 things, to declare a first and last name variable as well as call the function we create below that one. In the decloration of the function we create you will notice the word string next to F, and L. This tells the compiler which dataype the function expects. If we call this function with a variable other then a string we can get an error. The function call and decloration of the variables are the only real difference between this file and the Hello_V file. The Print statement uses verbs just like in Hello_V however they are declared and come from a place outside of the function it is called.

## Hello_Pointer, Hello_Hijack, Hello_Proper
Because pointers are going to be refrenced in a later section this will not be gone over in to much detail. However while googling how to pass strings to a function in Go I noticed others metioned using a pointer, and therefore I wanted to make note of them and using them in a way that does not result in data overwriting. This is because pointers are the location in memory for the given variable and not the data the variable holds. These 3 files indicate how pointers are passed to a function "Using & to represent the address, and using the * to derefrence the pointer passed in."

If you are experianced in C or other low level programming languages you may be familur with pointers. If you are new to pointers dont worry to much if you dont understand whats going on just yet. Just know that if you modify the value of a derefrenced pointer anywhere in the code it is passed to, then it will be modified everywhere in the code.