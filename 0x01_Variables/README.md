# 0x01_Variables
### Dataype Time
Quick overview:
	This directory is dedicated to datatypes, below I will outline the datatypes used in the files above. In the table I will give a brief overview of the general datatypes that appear to be present in most programming langauges, however because Go is compiled we have some more flexability then say python when it comes to saving/or storing data. This means that Go will require a little more in-depth knowledge about these dataypes, and we will go over those specifics as we get to them in the examples.

## Data Types:

| Dataype   | Description |
| :-:       | :---        |
| "String"  | A string is a representation of a collection of Letters and Numbers. Some examples include "Bobby", "ASDFUIAHE", "Happy Birthday Tony". They are most often denoted with quotes. |
| "Int"     | Int stands for Integar. It is the dataype that represents whole	numbers. The reason I say whole numbers is because an int can only be a whole number, there is a completly different datatype for numbers with a decimal point. For example "1024" |   
| "Float"   | A float is the datatype that represents numbers with a decimal point. For example "3.14" |
| "Boolean" | A Boolean is the representation of True or False. The Boolean datatype will consist of those 2 predefined constants |

## Variables
Compaired to the files in the last directory hello_world, it would appear as there is a lot more code in this file, while there is some truth to that there really is only one or two new concepts.

Declaring a variable will work pretty much the same way for every datatype in this directory, we do that by using the key word ``` var ```. This keyword is telling the compiler, "Hey, this next word is going to be important, it'll represent some data. Hold on to it so when I need that data I can refrence it.". This means whatever word, character, or series of characters(starting with a letter) comes after "var", will be the "variable name". After the variable name, we tell the compiler what type of data this variable will store. This is where datatypes come into play.

## 00_Strings
A string is any sequance of characters understandable by the computer, usually in the form of Key Words(like "apple" or "Save"), names, sentences, address, etc. There are many reasons we would want to modify strings, however this can become a more challanging topic depending on what you are after so for this file we will keep it fairly light.

In this example strings are printed and combined(also known as concatenated). Here are some key take aways from what the example is demonstrating.

- The keyword used when creating a string variable is "string". ``` var MyName string = "Cody"```

- Each string we declare will need to be enclosed in quotes, this will tell the compiler that we want to have everything in between them saved.

- By using a plus sign, we can "Combined", or append string b to string a

- We can store strings as a variable, and then combined them with raw strings, if need be.

Directories down the line, for a lack of better terms, will cover using strings in a more complex manor. That being said as long as we have a basic understanding of strings we should be fine for now moving forward. If you are new to programming and want to challange yourself a little bit, try to save the string ``` Sure, Monday's are the "best" day of the week ``` into a variable with the name "Monday". To help you out a little bit, this should lead you to articles on "escaping" a character in a string.

## 01_Ints

### **Why so many???**
This was probably your first thought when you saw all the variable declorations in this example. Lets take a minute to understand what we see here. The best way to do that is if we review the output of this file together.
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

Above we see some single character variables, and a number assigned to it, and then we see those same variables, but with a datatype assigned to it. Each one of these variables is a type of int (think of the word integar), so why do we have an int8, or an int64?

Those numbers are a representation of the size in bytes the number is allocated in memory. Back in the day computer programmers had to account for every bit of data they wanted to store, create, or work with, and back then. Storage capacity wasn't even close to where it is today. This meant that to keep things small, they assigned a normal int to 4 bytes worth of data. If you wanted more storage space you had to ask the compiler to store the data in a "long int" which is 8 bytes worth of data.

Now you may be asking yourself why you should care about how much data a variable for a number stores. Unfortunatly to answer to that question would take a very long time, so I'll sum it up. It all comes down to the size of the actual number, the compiler will take a number like 1 and turns it into binary, which is 0001 in 1 byte binary. When dealing with 4 bytes, we can only store a number as large as 2,147,483,647, or -2,147,483,648. Once we get to any number after that, in order for the computer to be able to know what it is, it will need a 5th byte to store the extra values in.

Okay one more queastion you may be asking yourself. Why does a modern day programming langauge have us worry about things like bytes? While it can be a boring topic to learn it is important to have an understanding of where space is coming from if you are trying to make a small, compact program that can be ran at the least amount of expance on the machine. While it is not super important to save space on a program that adds all your expences together, it will be important if you are working for a company that needs a task completed on lets say a server, that has very limited resources to expense.

If you are intrested in learning more on this topic, you should check out some of the links below.

https://tour.golang.org/basics/11

https://www.tutorialspoint.com/go/go_data_types.htm

https://docs.microsoft.com/en-us/cpp/cpp/data-type-ranges?view=msvc-160  

## 02_Floats


## 03_Booleans


## 04_Complex