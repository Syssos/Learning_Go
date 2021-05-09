# Functions

## Importance

When making more complex programs, repeating tasks, or large chunks of code can be a common accurance. Meaning a method of easily re-using code can be vital to a programs readablity, keeping it as easy for the developer to follow along as can be.

This is where functions come in, this directory will cover the basics of functions, and some **import**ant uses of them.

## Function Comments

In the examples above we look at a new way of writing comments. This style of comment is known as a style block. This is the point in which we start using functions in our examples, and to keep things organized these comments will be used from here on out. The following code block will demenstrate how the comments are layed out, its fairly simple.
```go
/*
	Description of what function acheives
	Return: what dataype is returned / a description 

*/
```


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


## 02_Return

Return is one of the key fetures of a function that makes them useful. With a a return statement you can make a function that takes some data, and turns it into a new set of data. The name explains it pretty well, it will return the data to the function call.

In the example we see that the function returns a value, which is then used as a number in a math problem. That result is then saved in a variable that is printed.

Go offers us the ability to return multiple values, in the example we see that there is the float keyword next to the function call, this tells us that the function will return a value of the type float. If we wanted to return more values, we would have to start at the function decloration. 

Below is an example of what that function decloration would look like.
```go
func return_test(x int, y int) (int,int) {...}
```

In this example theres 2 sets of paranthesis. This is because the second set houses the datatypes for each value returned.

```go
return value, err
```
above is an example of what the return statement would look like.

A good use case for this is demenstrated in 01_Imports, when we set the function to a variable lang and err. If there was an issue in that function the error value would not be nil, and we would be able to print what the error message is. This is a good way to handle issues that could arise in a program.


## Conclusion

When working with any kind of programming language functions are always yout friend. Keeping code clean and easy to read is the best way to keep track of whats going on. Functions not only allow us to acheive this, but they give us a way to use the same "Chunks" of code, whenever we may need it. Understanding how functions and returns work is extremly important and it is highly recommended that you do further research if you still have questions on how they work.

[Back To Top](#functions)