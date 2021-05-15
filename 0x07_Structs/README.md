# Structs

Structs, or structures in Go are very handy little methods of clumping together different types of data. With structs we can get some functionality higher level, Object-Oriented, programming languages have to offer. While Go supports object-oriented styles, there is no hiearchy in place making it an object oriented languages.

Through the struct examples we will be looking at a "car" as an example. Due to the nature of a car this will let us use a combonation of datatypes.

## 00_Structs

In this example we go over how structs can be utilized to hold different types of data. To start things off you may have noticed that the structure decloration is outside of the main function. Structures are declared out

## 01_Adding_Values

## 02_Removing_Values

Before you start saying that we aren't removing values were setting them to 0, I just want to say, thats exactly what's happening in this example. The reason for that is because we cannot actually remove the "Ratings" attribute from the struct, without saving the data into another struct that doesn't contain it.

This is because at runtime the structs type definition cannot be changed. This means there is no way for the code to "update" the values a struct holds on the fly. If a struct is not supposed to contain the data the only way to essentually remove it is by setting it to "empty".

## 03_Methods