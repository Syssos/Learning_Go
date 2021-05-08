# Functions

## Importance

When making more complex programs, repeating tasks, or large chunks of code can be a common accurance. Meaning a method of easily re-using code can be vital to a programs readablity, keeping it as easy for the developer to follow along as can be.

This is where functions come in, this directory will cover the basics of functions, and some **import**ant uses of them.

## 00_Functions

In this example we look at declaring, and calling a function. We declare a function by using the Go keyword "func".

The words following func is a name of our choice and it is how we will refrence it in our code. The Paranthese that follows the name of are required by Go, it is essentually a list of arguments the function will take. If our function takes no arguments then we can specify that by leaving it blank.

We declare multiple functions in this example but note that the main function is still part of the code, and is the entry point in which we will begin.

## 01_Imports

While creating functions or a collection of functions is really our end result as devolopers, Go gives us the option to utilize functions in a way that set's it apart. Like most laguages Go offers built in librarys we can take advantage of, for example "fmt". We've been importing it this whole time so we can print stuff to the screen. 

As we see in the example there are also built in functions like math, which gave us the Abs function to find the absolute value of our "number" variable. Its recommended that you spend some time researching the Go's "Standard Libraries" as there are plenty of functions available, and some could be of use to you.

Aside from built in libraries this example also outlines the use of a function from a package created by [endeveit](https://github.com/endeveit "endeviet home repo") here on github. This is a neat little go package that guesses the langauge a string is written in, such as english or spanish. 

To use this function I had to utilize the following Go command.

```
go get -u github.com/endeveit/guesslanguage
```

This command will search for the package we are after, download it, and place it in our $GOPATH so we have access to it when we need it. As the example shows when using a package from github, another online repo or source we do not us the HTTPS:// that a browser would require as the compiler does this for us. This is a nice touch as it allows us to keep our code clean and concise.

After the command above is ran, the import command shouldn't have any problems pulling in the function we are after, below is how we would call the function if we were not using a variable in place of the string, in both cases the output should be "en" as the sentences are in english.

```go
guesslanguage.Guess("This should output english, at least I think it will, I and it may do it.")
```

Hopefully this outlines the power Go provides developers when it comes to shareing code, and creating more and more powerful programs. If we were to write a program that took some text and translated it into another langauge, a package like this could be utilized to cut back on the amount of code we would have to write in order to make it happen.

## 02_Some_Standard_lib

In the last example we saw how to import packages from different sources, two from the standard library Go offers, and one from github. While the import from github shows us a door to infiate possibilites, lets slow down a bit and take a look at what Go bundles up and gives us at the door.

```go
import (
	"strings"
	"math"
	"fmt"
)
```
<!-- strings.ContainsAny
	 strings.ToUpper("ThIs Is So CoOl Of Go!") -->
