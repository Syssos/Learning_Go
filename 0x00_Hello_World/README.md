# Hello_World

## Gettings started 

Each one of the files above has a similar job, print some data. How we do that is different in each file. The goal here is to have an understanding of printing data to the screen before moving on. We will be using print functions all over the place, to see what the code is doing when being run.

Okay so how do we get started?

In Go we work inside of a package. This will be explained later on down the line, but for now all we need to know is that means code will always have to start with the same line.

```go
package main
```
Every Go script in this example will work within the main package. The scripts we will start off writing are by no means complex and will usually work within one main function.

An “import” statement, to keep things simple, is basically telling the compiler "hey, we need this chunk of code, can you go find it so I can use some of the code within it?". In the case above we are calling the "Format" package. This is a built in package and is what will give us the ability to print strings to the screen.

```go
import "fmt"
```

A function declaration will always start with the "func" keyword. We will go into these in more depth at a later point in the repository, for now there is only one key takeaway we need to have. The Main function in Go is our starting point. When a Go script from the examples to come is run/or executed the first thing it will do, to keep it light, is call the main function.

```go
func main() {...}
```

Every Go program needs to start within the main function.

## Hello_World [</>](https://github.com/Syssos/Learning_Go/blob/main/0x00_Hello_World/00_Hello_World.go)
Now that we have a general Idea of the basics a Go script needs, let's get into writing some code. In this file we see the package, import, and main function declaration lines of code mentioned above, however we see something in the main function that we didn't go over yet, so let's do that now.

```go
func main() {
	fmt.Println("Hello World!")
}
```
Within the main function we see the word "fmt" with the word "Println" attached to it with a period. What does this mean?

Println is a function that prints a string we pass it (everything in the parenthesis next to it), in order to use that function we need to tell the compiler where it is. We do that by adding the "fmt" in front of it, connected with a period.

This is essentially saying, in fmt, there is a function called Println.

An even shorter way of saying it

in fmt, theres a function, Println

or even shorter

in fmt, Println

how a computer sees it

fmt.Println

## Verbs [</>](https://github.com/Syssos/Learning_Go/blob/main/0x00_Hello_World/01_Verbs.go)
In this file we see 2 new things, the first one is a new print function, Printf, the second is a "%v" in one of the arguments we hand it.

The reason we went with a new print function from fmt, is so we had access to Verbs. A Verb in Go is essentially a placeholder within a string, that will then be later replaced with some kind of value. In the example we see that there are 2 sets of strings in quotes, separated by a comma being passed to the printf function. The first being "%v World!", and the second being "Hello". The first string is the string we want the function to print, the second is the value we want to replace the "%v" with.

All said and done the printf function will return a string that reads, "Hello World!". If we replaced the second string, with lets say "Oh My", the main function will print the string

"Oh My World!"

## Double_Verbs [</>](https://github.com/Syssos/Learning_Go/blob/main/0x00_Hello_World/02_Double_Verbs.go)
In 01_Verbs we saw that we can use %v to insert one string into another. What if we wanted more, so we could do things like print a greeting. This example shows how that can be achieved.

When using the printf function we can use as many verbs as we'd like, however we need to make sure we pass the same number of arguments. Too little or to many arguments can result in an error, an easier way to look at what is going on is if we use "sudo code" or "Fake code" to better explain what is happening.

```
("hello <variable_1> <variable_2>, How are you?", variable_1, variable_2)
```
Above is a line of code that wouldn't exactly work if we wrote it in the Go script, but it helps breakdown what is happening a little better, there are 3 arguments being passed to printf, the first is the string we want to print, and as we can see, each value after that will correspond with a spot in the string we want to print.

That being said we could make an entire string out of verbs. Although it would look pretty weird.
```go
fmt.Printf("%v %v %v, %v %v %v", "Hello", "Thomas", "Edison", "How", "are", "you?")
```
this will give us the same output as the example in the file.

## Conclusion
The major point of starting with this topic first is so we have a way to see what we are doing with the code. In the examples to come we will be doing all sorts of things with words and numbers. We have to have some way to see it. The Printf and Println functions are going to be used in a majority of the examples you see throughout the repository so make sure you spend some time playing with them, and even doing some further research on them if need be.

A good milestone to know if you're ready to move on is, from scratch without external help, create a go script that prints a sentence, and run it successfully. If you can do that then you shouldn't have any issues with what's to come.

[Back To Top](#hello_world)
