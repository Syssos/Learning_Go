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

While loops are another form of loop that works without declaring the count variable or how it increments. Instead a while loop will loop while what in the condition section evaluates to true.

In the example we see that the count variable is incremented as the last thing inside the brackets. This is because the for loops decloration no longer handles that task. If we did not increment the count ourselves, outside of the for decloration, then we would have a loop that runs indefinetly as the condition will always result in true. Some cases will require loops that run indefinetly, or until the condition is changed to say otherwise.

## 02_Arrays

Arrays in Go are datatypes, that allow us to store a collection of other types of data, such as ints or strings. They are handy when you have something like a list of names or id numbers you need to save.

If you are used to other programming languages like python, then you may be familiar with the list dataype as a way to represent a collection of keyless data. Go doesn't use lists but instead arrays are used to store data in this manor. If you are not familiar with arrays lets take a look at how the computer will veiw them.

Picture the following array
```go
arr := [4]{ "Go", "is", "super", "fun"}
```

| Memory location   | Data Stored in memory |
| :-:               | :---                  |
| 0x8130            | "Go"                  |
| 0x8131            | "is"                  |   
| 0x8132            | "super"               |
| 0x8133            | "fun"                 |


Above is a chart that represents roughly how a computer will store an array. The addresses on the left are addresses in memory, the data on the right represent data stored at those addresses. The reason we specify the length of the array is because it tells the compiler to take x amount of addresses in memory and reserve them to store data.

This allows us to pull data out of a list based off index. Lets say we want the word 'super' the third word from the list, arr[2] will give us this because arrays start at 0. When the compiler see's arr[2] it will start at the first location in memory for the arr array(0x8130) and then move forward until it gets to the 3rd address in memory that is tied to that array(0x8132).

## 03_Arrays_and_Functions

There will be times were we need some work done on an array, and want to create a function to handle it for us. This example shows how we can pass an array to a function. Simliar to how we pass other dataypes to a function we must declare a variable name and a datatype. In this example we see 2 different function calls, one using a set length arrays as arguments and another using no set length arrays

If we are working with an array that has a set size and need to pass it to a function, then the function will need to accept an array of that size. If you test this with an lengths that don't match you will see that you get an error from the compiler for the wrong datatype being passed to that function.

We also see that the data returned is an int. While for what we are after in the code this is perfect what if we want to return an array?

```
func Array_Math_noSize(arr3 []int, arr4 []int, a int, b int) []int {

	new_arr := []int{1,2,3,4}
	return new_arr
}
```
The code above is an example of returning a new array instead of an int. Note the change in the function call, at the very end, the int now has brackets infront of it, "[]int". This tells the compiler that the function will return an unsized int array.

The code above will have errors if you try to run it, this is because the arguments are declared and never used. However it is recommended that if your fuzzy on this topic at all that you play around with it. Google every error you run in to and try to get a sense of what the compiler will want or is asking for.

## 04_Arrays_in_range

Lets recap 02_Arrays breifly. Arrays are stored in memory with an address to keep track of the data. This must mean there is a start and end to the array.

If we were to want to do "work" on every object on an array we can use the range built in. This example shows us using range, to itterate through the list in a better fasion. If you are used to python or are familiar with it, the range built in gives us somewhat simliar uses as a for in loop.
```python
	for item in list:
``` 
Even though the code snippet above is in python is outlines what the forloop in this example is doing.

Think of how we itterated through arrays in previous examples. We used a for loop with a count, and set the forloop to stop at the same time the array would. While this works it can lead to errors. Range fixes that for us by handling the loop count, based off array length. This is why we do not see a counter being set or incremented, and why it is similar to a for in statement in python. It will do "work" for every item in the array.

## Conclusion

This directory outlined for loops and arrays, and hopefully shows the two can go hand in hand. If you want to work with each peice of data in an array, a for loop is a clean and concise way of acheiving it. Aside from arrays for loops will also have a large amount of use cases, weather it be a web application or command line based scripts for loops can be found all over the place. That goes for arrays as well. If any of these topics are still something that you have questions on you should deffinetly continue researching.

A great way to continue practicing these topic's would be to combine these, and everything we've made so far, and make a more complex program such as a game like rock paper siccors. Don't be afraid to research topics and do some further reading into errors.

[Back To Top](#for-loops-and-arrays)