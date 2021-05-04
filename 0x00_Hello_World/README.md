# 0x00_Hello_World
## Gettings started 

Each one of the files above has a simlilar job, print some data. How we do that is different in each file. The goal here is to have a solid understanding of printing data to the screen before movinbg on. This is because printing data to the screen(or the "standard output") is one of the most important features in most programming langauges. We will use print functions all over the place, and can they be your best friend while debugging issues in more complex scripts.

Okay so we have an idea of what we want our code to do,  but how do we get started? Well lets start at the top of the page and work our way down.

Each one of the files above will start with the same 3 lines so lets review those lines to have a better understanding.

```go
package main
```
The line above is a manditory line in Go. This is because every Go script will start within the main package. The scripts we will start off writing are by no means complex enough to have real value in importing into another project, so for now we will just need to know that every file will start with this line. When the Repository is finished and all the steps are added, I will create an example project demonstrating and explaining in greater detail how to create a module and export it in a way others can utilize.

```go
import "fmt"
```

An import statement, to keep things simple, is basically telling the compiler "hey, we need this package of code, can you go find it so I can use some of the code within it?". In the case above we are calling the "Format" package. This is a built in package and is what will give us the ability to print strings to the screen.

```go
func main() {...}
```

The line above is a function decloration. We will go into these in more depth at a later point in time, for now there is only one key take away we need to have. The Main function in Go is our starting point. When a go script is ran/or executed the first thing it will do, to keep it light, is call the main function.

Every Go program needs to start within the main function.

## 00_Hello_World
Now that we have a general Idea of the basics a Go script needs, lets get into writing some code. In this file we see the package, import, and main lines of code mentioned above, however we see something in the main function that we didn't go over yet, so lets do that now.
```go
func main() {
	fmt.Println("Hello World!")
}
```
Within the main function we see the word "fmt" with the word "Println" attached to it with a period. What does this mean?

Println is a function that prints a string we pass it (everything in the paranthesis next to it), in order to use that function we need to tell the compiler where it is. We do that by adding the "fmt" in front of it, connected with a period.

This is essentially saying, in fmt, there is a function called Println.

An even shorter way of saying it

in fmt, theres a function, Println

or even shorter

in fmt, Println

how a computer sees it

fmt.Println

## 01_Verbs
In this file we see 2 new things, the first one is a new print function, Printf, and a "%v" in one of the arguments we hand it.

The reason we went with a new print function from fmt, is so we had access to Verbs. A Verb in Go is essentually a placeholder within a string, that will then be later replaced with some kind of value. In the example we see that there are 2 sets of strings in quotes, seperated by a comma being passed to the printf function. The first being "%v World!", and the second being "Hello". The first string is the string we want the function to print, the second is the value we want to replace the "%v" with.

All said and done the printf function will return a string that reads, "Hello World!". If we replaced the second string, with lets say "Oh My", the function will return the string "Oh My World!"

## 02_Double_Verbs
In 01_Verbs we saw that we can use %v to insert one string into another. What if we wanted more, so we could do things like print a greeting. This example shows how that can be acheived.

When using the printf function we can use as many verbs as we'd like, however we need to make sure we pass the same number of arguments, to little or to many arguments can result in an error, an easier way to look at what is going on is if we use "sudo code" or "Fake code" to better explain what is happening
```
("hello <variable_1> <variable_2>, How are you?", variable_1, variable_2)
```
Above is a line of code that wouldn't exactly work if we wrote it in the Go script, but it helps breakdown what is happening a little better, there are 3 arguments being passed to printf, the first is the string we want to print, and as we can see, each value afterthat will correspond with a spot in the string we want to print.

That being said we could make an entire string out of verbs. Although it would look pretty weird.
```go
fmt.Printf("%v %v %v, %v %v %v", "Hello", "Thomas", "Edison", "How", "are", "you?")
```
this will give us the same output as the example in the file.