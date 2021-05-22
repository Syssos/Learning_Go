# Unit Tests

## Overview

Unit tests can be extremely helpful for making sure code functionality doesn't break while adding or changing features in a program. Testing code is a common thing for programmers to want to do, so the developers of Go have created a way to do that with a built in package known as "testing". The testing package will contain a majority of the tools we will need for testing our code. This directory will overview some of those tools so you can get to testing your code.

## add [</>](https://github.com/Syssos/Learning_Go/blob/main/0x10_Unit_Tests/00_add.go)

This example is a pretty short one, containing a function, Sum, and the main function. We can see that the Sum function takes 2 int arguments and will return an int. The main function is responsible for printing the result of the Sum function to the screen.

While this may seem straight forward there is one thing I would like to point out before moving on. That is the sum function will return a value. This is important with the type of testing, the "testing" package allows for. 

When using the "testing" package, the main objective is to ensure a specific function is returning values as it should. Testing a function that prints to screen, or doesn't return values is not nessesairally desired as there would be a different reason as to why a user wouldn't see output to the screen, which also could mean the code is unable to build. 

Usually making sure the output that the user would see like a returned value is more important. Which is where these tests come into play.

## add_test [</>](https://github.com/Syssos/Learning_Go/blob/main/0x10_Unit_Tests/01_add_test.go)

When creating a test file we want to name it ending in "\_test". The reason for that is because when working within a module's or package's folder, we can run ``` go test . ``` and all of the files ending with the test suffix will run. If no errors are found then we should see an output of ``` ok      github.com/Syssos/unit_t        0.001s ```.

Lets now take a look at some things that are happening in this example.

To start off we need to import the "errors" package, as this will contain some useful tools for use to use.

Next we start declaring a function with a name that starts with "Test", this is to signify that the function is used as part of a "Unit Test", and to make it exportable. We then declare an argument of type ``` *testing.T ``` to be taken by the function.

When running go test, go will call this function and pass it an argument of a testing struct. This is what gives us the ability to flag an Error or Failure, via t.Error, t.Fail.

We see we can do that in this example when a desired result is not returned from the function call.

To test for the desired result we essentually created a slice of structs, with the name cases. In each struct instance was the x value, y value and the desired result. We used those values to test for correct output.

If you come across certain things that you would like to log, but not flag as an error, you can do so via ``` t.Log() ```. This will log the event but not trigger an error or failed tests.

You may also want to add more details like values to the error message. This can be done simular to how you print values with printf, however in the case of an error you would use ``` t.Errorf() ```

## Conclusion

It is highly recommended that further research is done into unit testing. The topic is very broad and each function can have different cases, datatypes, and results. However making sure that code is tested and can be tested after every major/minor change will allow for you to write more effiecent code as you will see if anything is broken or is being caused to break in ways that can cause for the program to work incorrectly. To learn more about unit tests or the "testing" package I would check out the official [Golang](https://golang.org/pkg/testing/) site and read up on what they have to say about testing.