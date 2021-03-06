# Variables

This directory is dedicated to datatypes, below I will outline the datatypes used in the files above. In the table I will give a brief overview of the general data types that appear to be present in most programming languages, however because Go is compiled we have some more flexibility then say python when it comes to saving/or storing data. This means that Go will require a little more in-depth knowledge about these data types, and we will go over those specifics as we get to them in the examples.

## Data Types:

| Dataype   | Description |
| :-:       | :---        |
| "String"  | A string is a representation of a collection of Letters and Numbers. Some examples include "Bobby", "ASDFUIAHE", "Happy Birthday Tony". They are most often denoted with quotes. |
| "Int"     | Int stands for Integar. It is the dataype that represents whole	numbers. The reason I say whole numbers is because an int can only be a whole number, there is a completly different datatype for numbers with a decimal point. For example "1024" |   
| "Float"   | A float is the datatype that represents numbers with a decimal point. For example "3.14" |
| "Boolean" | A Boolean is the representation of True or False. The Boolean datatype will consist of those 2 predefined constants |

## Variables 
Compared to the files in the last directory hello_world, it would appear as there is a lot more code in this file, while there is some truth to that there really is only one or two new concepts.

Declaring a variable will work pretty much the same way for every datatype in this directory, we do that by using the keyword ``` var ```. This keyword is telling the compiler, "Hey, this next word is going to be important, it'll represent some data. Hold on to it so when I need that data I can reference it.". This means whatever word, character, or series of characters(starting with a letter) comes after "var", will be the "variable name". After the variable name, we tell the compiler what type of data this variable will store. This is where datatypes come into play.

## Strings [</>](https://github.com/Syssos/Learning_Go/blob/main/0x01_Variables/00_Strings.go)
A string is any sequence of characters understandable by the computer, usually in the form of Keywords(like "apple" or "Save"), names, sentences, address, etc. There are many reasons we would want to modify strings, however this can become a more challenging topic depending on what you are after so for this file we will keep it fairly light.

In this example strings are printed and combined(also known as concatenated). Here are some key takeaways from what the example is demonstrating.

- The keyword used when creating a string variable is "string". ``` var MyName string = "Cody"```

- Each string we declare will need to be enclosed in quotes, this will tell the compiler that we want to have everything in between them saved.

- By using a plus sign, we can "Combined", or append string b to string a

- We can store strings as a variable, and then combine them with raw strings, if need be.

We will cover using strings in a more complex manner later on. As long as you have a basic understanding of strings we should be fine for now moving forward. If you are new to programming and want to challenge yourself a little bit, try to save the string ``` Sure, Monday's are the "best" day of the week ``` into a variable with the name "Monday". To help you out a little bit, if you do research on this, it should lead you to articles on "escaping" a character in a string.

## Ints [</>](https://github.com/Syssos/Learning_Go/blob/main/0x01_Variables/01_Ints.go)

**Why so many???**

This was probably your first thought when you saw all the variable declarations in this example. Let's take a minute to understand what we see here. The best way to do that is if we review the output of this file together.
```
c: 127,
o: 1568,
d: 654351,
e: 35138484384,
h: 255,
a: 2563,
ck: 532653,
s: 313843543

c: int8,
o: int16,
d: int32,
e: int64,
h: uint8,
a: uint16,
ck: uint32,
s: uint64
```

Above we see some single character variables, and a number assigned to it, and then we see those same variables, but with a data type assigned to it. Each one of these variables is a type of int (think of the word integer), so why do we have an int8, or an int64?

Those numbers are a representation of the size in bytes the number is allocated in memory. Back in the day computer programmers had to account for every bit of data they wanted to store, create, or work with, and back then. Storage capacity wasn't even close to where it is today. This meant that to keep things small, they assigned a normal int to 4 bytes worth of data. If you wanted more storage space you had to ask the compiler to store the data in a "long int" which is 8 bytes worth of data.

Now you may be asking yourself why you should care about how much data a variable for a number stores. It all comes down to the size of the actual number we want to store. The compiler will take a number like 1 and turn it into binary, which is 0001 in 1 byte binary. When dealing with 4 bytes, we can only store a number as large as 2,147,483,647, or -2,147,483,648. Once we get to any number after that, in order for the computer to be able to know what it is, it will need a 5th byte to store the extra values in.

Okay one more question you may be asking yourself. Why does a modern day programming language have us worry about things like bytes? While it can be a boring topic to learn, it is important to have an understanding of where space is coming from if you are trying to make a small, compact program that can be run at the least amount of expense on the machine. While it is not super important to save space on a program that adds all your bills together, it will be important if you are working for a company that needs a task automated on lets say a server, that has very limited resources to expense.

If you are interested in learning more on this topic, you should check out some of the links below.

https://tour.golang.org/basics/11

https://www.tutorialspoint.com/go/go_data_types.htm

https://docs.microsoft.com/en-us/cpp/cpp/data-type-ranges?view=msvc-160  

## Floats [</>](https://github.com/Syssos/Learning_Go/blob/main/0x01_Variables/02_Floats.go)
Like the int data type, the number after the keyword float is in reference to the amount of storage space allocated to the variable. Due to the nature of floats, more space is required to save everything before and after the decimal point, meaning the amount of bytes needed to represent a number is greater. This is why we do not see a float8 or float16 like we did with integers.

Floats are composed of ints, so let's say we have a float32, let's break that down in a more understandable way

3.14

lets see the minimum requirements to save this number

[1 byte].[1 byte]

As you can see the total number of bytes is only 2, which would make up a float16, however that only allows us to store a number as large as 127.127, and there's many cases in which this will simply not work.

Let's make it a little more useful by adding 1 more byte on each side of the decimal.

[2 bytes].[2 bytes]

More bytes means we can store bigger numbers. By adding a byte to each side of the period we can now store a number as large as 32767.32767, and the float variable will be 4 bytes in size, or a float32

While this is not exactly what's going on behind the scenes, it should give you a better idea as to why we only see 2 float types opposed to the amount of int types.

Below is a chart for float sizes

| Dataype   | Bytes |            Size          |
| :-:       | :---  | :---                     |
| "float32" |   4   | 3.4E +/- 38 (7 digits)   |
| "float64" |   8   | 1.7E +/- 308 (15 digits) |

## Booleans [</>](https://github.com/Syssos/Learning_Go/blob/main/0x01_Variables/03_Booleans.go)

Booleans are a rather simple datatype as it is a representation of true or false. They are important to touch base on however because the relevance of true or false in programming is pretty vast. As you will see in this example the "if" statement is checking **if the variable is true or false**, if it is true it will run the code in the brackets following. If it is false it will move on to the next part of the code.

The concept of true or false will be seen all over the place in Go code. While you may not always be setting a variable to true or false, comparisons that guide the code's functionality will heavily rely on this concept.

## Complex [</>](https://github.com/Syssos/Learning_Go/blob/main/0x01_Variables/04_complex.go)

Complex numbers in Go are not something you will need to worry about when writing day to day Go scripts. However it is a concept we should touch on so you're aware of their presence in the language.

To explain complex numbers let's look at how the computer stores them.

``` complex64 = (<float32>, <float32>) ```

or

``` complex128 = (<float64>, <float64>) ```

As we see above a complex number is essentially 2 float variables combined. However when storing a complex number the 2nd argument is seen as an imaginary number. The example for complex numbers outlines that, and also shows us some basic Arithmetic between 2 complex numbers. Please note that any kind of arithmetics will require complex numbers of the same type.

## Constants [</>](https://github.com/Syssos/Learning_Go/blob/main/0x01_Variables/05_constants.go)

Constants in Go are variables that are not designed to change. If you attempt to update a constant somewhere in the Go code after declaration, the compiler will complain. Constants can be a character, string, boolean, or numeric value. As you will see in a later example Go can choose the datatype for us, however constants need to be declared with the keyword const, so the go compiler knows this is a special value that is not intended to change.

## Type_casting [</>](https://github.com/Syssos/Learning_Go/blob/main/0x01_Variables/06_Type_casting.go)

All of the above outlines how important types are to Go, but what if we have a 42 as a string, and need to add it to a 7 that's an int? Well Go lets us do that with "type conversions". This example should look pretty straight forward if you've been following along. We do have to keep one thing in mind however, the data type conversion **must make sense**. If we try to store a value such as "S" into an int8 or int32 we will get an error, this is because a string can store non-numerical values, which cannot be stored as an int. Even with a string being only "42" the compiler will refuse as it is much more time consuming to skip this checking process, and Go is all about speed.

However the example does outline some working examples of type conversion. This will mainly be used when trying to do some kind of arithmetic calculations on numbers not stored as the same datatype.

## Composite [</>](https://github.com/Syssos/Learning_Go/blob/main/0x01_Variables/07_Composite.go)

Now that we have a basic understanding of the data types in Go, and know we can change variable types as long as it makes sense, let's look at one of the cool features Go guideve developers. For obvious reasons constantly declaring variable types can be hard to track, or time consuming to write out. Let's be honest, as developers we already do a lot of typing, so the less we have the better.

With the ":'' sign before the "=", we are telling the compiler to assign a data type to the variable, that makes sense, for us. If we say ``` x := 3.14 ``` the compiler will look at the data we want to save, see that it has a decimal point, and know to save that variable as a float. This is a much more convenient way of creating variables.

## Blank Identifiers

Before we conclude this section I would like to mention blank identifiers(```_```). While they may not be super useful at the moment, they are important to know about. As we expand our code we will run into times we are given more data then we need. Go doesn't allow for us to create variables and not use them, so in situations like this, using the blank identifier can save us. Below is an example using pseudocode.

If you declare 2 variables to store data returned by a function, and only want to use 1, the go compiler will give you an error saying dif was declared and never used.
```
sum, dif := SumDif(2, 3)

printf(sum)
```
Using a blank identifier will tell the compiler to ignore that value from the function and continue with no errors.

```
sum, _ := SumDif(2, 3)

printf(sum)
```

## Conclusion

Datatypes in Go will soon become a small part of the language as we learn and develop our skills, however it is important to know what is going on in the computer's brain to write efficient code. If you are new to programming, and are still a little fuzzy on any of the datatypes above, it is strongly recommended that you do further research on this topic, as it is not just Go that use's this method of storing information and understanding it is an important part to being a better programmer all around.

[Back To Top](#variables)
