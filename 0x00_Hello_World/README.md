# 0x00_Hello_World
### Gettings started with Go
Before I get into the description of the code you see above, and below, I would like to mention that this is not my first programming langauge, so some of the things I touch on are a little more advanced then basic string printing but because it fits in-with the catagory I would like to mention them here.

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
This file contains multiple new conscepts outside of just string printing, Some of those I will highlight brefiely so any person new to programming can try to follow along.

```go
Firstname := "James"
```
The First concept is variable creation, Here we see that I declare the string "James" to the variable Firstname. A string is a dataype and is basically just a collection of characters, like a word, or sentance. We do that with ":=". I will touch more on this in the datatypes section. For now just now its how we set variables.

```go
func Greeting_user(F string, L string) {...}
```
The Second concept is creating a new function. This is another concept that will be reviewed later on. This is simular to the main function in the sense of it being a collection of code that runs when called. The same way the main function is called, we can call a function we create from within main itself. This allows us to keep code inside of main condensed. We see that it takes to arguments, F and L, and they are both strings. This means when we call the function we can either pass it 2 strings or 2 variables of type string.
```go
Greet_user("James", "Dean")
```
Strings
```go
Greet_user(Firstname, Lastname)
```
variables
We can see the variables dont have quotes. Meaning the compiler will look for a varaible with that name rather then passing the strings as is.

Aside from the new topics the function actually that prints the string is the same as the function from Hello_V, just taking new arguments.

## Hello_Pointer, Hello_Hijack, Hello_Proper
If you are new to programming or looking at how to print a string, these files will not be needed. The point of these files were to show the use of pointers in Go. This topic will be touched on later on, the purpose of them being here is to show another method of passing strings through a function, and then demenstrate why it may, or may not, be useful in a project.

A pointer is a refrence to a variables location in memory. These are denoted with &, they are derefrenced with the *, they will apear in the code as shown below
```go
&Firstname
```
Above will give us the location in memory of the varaiable, not the value the varaible holds.

In order to access the information at the address in memory we need to derefrence that variable, or in leiman terms, tell the compiler "I see where your located, but what value does that location hold." We use the "*"
```go
*Fristname
```