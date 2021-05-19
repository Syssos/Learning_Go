# Unit Tests

## Overview

Unit tests can be extremely helpful for making sure code functionality doesn't break while adding or changing features in a program. Testing code is a common thing for programmers to want to do, so the developers of Go have created a way to do that with a built in package known as "testing". The testing package will contain a majority of the tools we will need for testing our code. This directory will overview some of those tools so you can get to testing your code.

## 00_add.go

This example is a pretty short one, containing a function, Sum, and the main function. We can see that the Sum function takes 2 int arguments and will return an int. The main function will print the result of the Sum function to the screen.

While this may seem straight forward there are some things I would like want to note before moving on. The sum function will return a value. This is important with the type of testing the "testing" package allows for. 

When using the "testing" package, the main objective is to ensure a specific function is returning values as it should. Testing a function that prints to screen is not nessesairally desired as there would be a different reason as to why a user wouldn't see output to the screen, which also could mean the code is unable to build. 

Usually making sure the output that the user would see like a value is more important. Which is where this tests come into play.

## 01_add_test.go

<!-- t.Log can be used to provide non-failing debug information -->
<!-- Calls t.Error or t.Fail to indicate a failure, call t.Errorf to provide more details -->
