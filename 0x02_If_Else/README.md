# If Else

## Introduction

In this directory we will talk about a very useful built in, in Go. That is the if statement. As you will learn, the word "if" is extremely powerful. It is a way for developers to control the flow of their code. It will be how most error handling is done, as it can be a tool to tell developers when certain conditions are not met. With this tool in your toolbelt your code will take a huge leap in functionality so make sure you understand these statements before moving on.

## If_Else [</>](https://github.com/Syssos/Learning_Go/blob/main/0x02_If_Else/00_If_else.go)

This example outlines what a basic if, else statement is achieved. It will look to see if a condition, or result of a condition is true or false. If the condition is true the code within the brackets following if is run. If the result is false then the code inside the else function is run, a pseudocode of an if statement makes it a little easier to understand.

```
if [condition] {
	code in here is ran when condition = true
} else {
	code in here is ran when condition = false
}
```
We use an emotion variable in this example, but if we look at its value in the comparison, we can see how the computer will get true as a result.

we use:``` if Emotion == "Happy" ```

because Emotion is "Happy", the compiler sees

``` if "Happy" == "Happy" ```

which results in true, causing the code following to run. Once the code in the if statement's brackets is run, the compiler will jump to the next line of code outside the rest of the if/else statement. This means that the code in the else statements brackets will not be run.

## Else_If [</>](https://github.com/Syssos/Learning_Go/blob/main/0x02_If_Else/01_Else_if.go)

In this example we see a very close program goal, print a message corresponding with an emotion. The use of if else statements give us the ability to add more emotions to check for. An if else statement works the same as an if, however it is used to further check code if the initial if statement is not true.

In most cases if, else won't be used in the fashion the example show's, why that will be seen in the next example. Let's look at a more practical use case for if else statements using pseudocode.

```go
if user_id > 1 {
	user exists and is signed in
} if else user_id == 0 {
	user is a guest session
} if else user_id == 1 {
	user is an admin
} else {
	there is an issue elsewhere in the code causing this variable to be set incorrectly
}
```
Above we can see that we are using the if else to check for a specific condition. Let's say this example is out of a web based application's backend code. This is what's running on the server hosting a website, and depending on the user_id will change what the user sees on their screen.

If the user_id is greater than 1 we know they must have been assigned an id and are signed in.

If the user_id is 0 we know some other code we wrote noticed them browsing as a guest and assigned it to 0, so other parts of the code will know the user is a guest.

If the user_id is 1, the same thing that happens as 0, however instead of limiting what the user sees, this will most likely increase what the user sees.

The else statement is being used as an error check in this case. If the user_id is "toaster", for example, we know that there is an issue elsewhere, causing the code to act incorrectly and need to fix that section.

## Switch [</>](https://github.com/Syssos/Learning_Go/blob/main/0x02_If_Else/02_Switch.go)

If, if, if it's not so. Switch cases are your best friend when in need of checking multiple conditions with one variable. This example for switch cases achieves the same end result as the previous example. However it achieves this in a way cleaner way.

The "switch" keyword followed by a variable or value will essentially tell the compiler there is a value "here" that needs to be checked.

We then use the "case" keyword to tell the compiler which conditions they must be checked against, each "case" acts the same as an if statement in a sense.

```
if switch_var == case_var {...}
```
It does this for each case, and if a case is never met, or results in true, the default statement is run.

Default for a switch statement will work the same way as the else statement.

## Recap [</>](https://github.com/Syssos/Learning_Go/blob/main/0x02_If_Else/03_Recap.go)

This recap highlights some very important concepts we covered in the last few directories, These are the absolute basics of the Go programming language. It's very likely that a majority of the topics discussed will be used in every script, or program you write in the future. Hopefully that outlines the importances of knowing these concepts. There is a large amount of online resources that can help you if there is anything I didn't explain in enough detail. I recommend spending some time writing scripts like this example if you are still a little fuzzy on the topic. Google each error message and over time the messages will start to make sense as you learn what the code is doing in the background.

[Back To Top](#if_else)
