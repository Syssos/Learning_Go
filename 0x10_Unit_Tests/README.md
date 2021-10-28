# Unit Tests

## Overview

Unit tests can be extremely an extremly helpful tool for ensuring code's functionality and performance. Due to testing code being common amongst programmers, the developers of Go have created a way to do that with the help of the "testing" package. While this package contains a majority of the tools needed for developers to test their code, it still leaves test cases, and validating output up to the developer. This repository is targeted at explaining how these tools can be used to create efficient ways of checking code.

## add [</>](https://github.com/Syssos/Learning_Go/blob/main/0x10_Unit_Tests/00_add.go)

This file contains the code that will be tested in against in the unit tests. Due to how the Golang testing library is configured, we do not have to do anything special to account for the tests when initial writing the code.

## add_test [</>](https://github.com/Syssos/Learning_Go/blob/main/0x10_Unit_Tests/01_add_test.go)

Test file names are denoted by the suffix <b>"_test"</b>. This tells the Go compiler that there are test cases inside of the file to be used while testing. Inside this file you will notice 2 "types" of functions, one starting with the prefix <b>"Test"</b>, the other with the prefix <b>"Benchmark"</b>.

<!-- When creating a test file we want to name it ending in "\_test". The reason for that is because when working within a module or package's folder, we can run ``` go test . ``` and all of the files ending with the test suffix will run. If no errors are found then we should see an output of ``` ok      github.com/Syssos/unit_t        0.001s ```.

Let's now take a look at some things that are happening in this example.

To start off we need to import the "errors" package, as this will contain some useful tools for use.

Next we start declaring a function with a name that starts with "Test", this is to signify that the function is used as part of a "Unit Test", and to make it exportable. We then declare an argument of type ``` *testing.T ``` to be taken by the function.

When running go test, go will call this function and pass it an argument of a testing struct. This is what gives us the ability to flag an Error or Failure, via t.Error, t.Fail.

We see we can do that in this example when a desired result is not returned from the function call.

To test for the desired result we essentially created a slice of structs, with the name cases. In each struct instance was the x value, y value and the desired result. We used those values to test for correct output.

If you come across certain things that you would like to log, but not flag as an error, you can do so via ``` t.Log() ```. This will log the event but not trigger an error or failed tests.

You may also want to add more details like values to the error message. This can be done similar to how you print values with printf, however in the case of an error you would use ``` t.Errorf() ``` -->

## Running Tests

## Conclusion

It is highly recommended that further research is done into unit testing. The topic is very broad and each function can have different cases, data types, and results. However making sure that code is tested and can be tested after every major/minor change will allow for you to write more efficient code as you will see if anything is broken or is being caused to break in ways that can cause for the program to work incorrectly. To learn more about unit tests or the "testing" package I would check out the official [Golang](https://golang.org/pkg/testing/) site and read up on what they have to say about testing.
