# For loops and Arrays

## Overview
In this directory we will see how we can use for loops and arrays to our advantage. For loops themselfs can be a simple concept, however depending on how there used they can be an incrediably powerful asset to the code you are writing. They are tasked to do the same chunk of code, until the condition is no longer true.

Unlike if statements there is no code that runs if the condition is false, so if a counter is not reset, or is out of place before a loop is ran, it can cause it to run incorrectly, if at all.

Arrays are also very useful. They group together chunks of data under one variable name, this could be a list a names, or maybe a list of credit card numbers. In this directory we will see how we can use them, and share them with other parts of our code.

## 00_For_Loop
A for loop in Go is essentually the only kind of loop that will really be used. This is because depending on how you use it you can get the functionality of multiple types of loops in other languages. This example outlines the classic use case of a for loop. In the for loop decloration we see something interesting happening so lets use sudo code to make it a little easier to understand.

```
for <counter being set>; <condition being tested each itteration>; <counter increment>
```
Above we can see what each section of the statement is doing. Right after the word for, we create the variable that will be used to track how many itterations of the loop we are doing.

The ";" is how we tell the compiler we are done setting that variable

After we set the variable we need to tell the compiler what condition we will check for before we run the code after each. In the example we are using the less then sign, and because we use a counter variable we say as long as the counter is less then 10 run the code. This means that if we have the condition ``` counter < 10 ``` the for loop will loop until counter is 10.

One thing the example outlines is when using the counter in the output, the counter started at 0. This shouldn't be a suprise as we set it to 0, we still got 10 messages even though the counter stopped at 9

This is because when the counter is at 10 the condition statement will look like...

```
10 < 10
```
This statement will result in false, because 10 is not less then 10 but equal.

## 01_While_but_For

## 02_Arrays

## 03_Arrays_and_Functions

## 04_Arrays_in_range

## Conclusion

[Back To Top](#for_loops_and_arrays)