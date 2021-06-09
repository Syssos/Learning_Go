# Structs

Structs, or structures in Go are very handy little methods of clumping together different types of data. With structs we can get some functionality higher level, Object-Oriented, programming languages have to offer. While Go supports object-oriented styles, there is no hierarchy in place making it an object oriented language.

Through the struct examples we will be looking at a "car" as an example. Due to the nature of a car this will let us use a combination of data types.

## Maps [</>](https://github.com/Syssos/Learning_Go/blob/main/0x07_Structs_and_Maps/00_Maps.go)

Maps in Go are a data type that allows for saving data in key value pairs. Maps are a built-in associative data type of Go's, in other languages you may see these referred to as hashes, hash tables, or dictionaries(dicts).

To create a Map we can utilize another built-in function known as [make()](https://www.godesignpatterns.com/2014/04/new-vs-make.html), or we can use composite literals. Note that make() is seperate from the built in new(). While these 2 seem like they do the same job, they do different things and return different data types.

Once the map is created we are now able to add values to it. Make() will initialize the map instance, and then return it for us to use. Meaning we can set a value to an index of "key" and it will save the key, value pair into the map, ie ``` mapname['key'] = "value" ```.

If we wanted to view the data in the map at index "key", we can reference it the same as setting the value, just without setting anything to it, ie ``` fmt.Println(mapname['key']) ```.

To remove values in a map we can use the built in delete() function. This will remove the key value pair from a map, based on arguments passed to it. The first argument is the map to remove the pair from, and the second argument is the key to the items to remove.

## Structs [</>](https://github.com/Syssos/Learning_Go/blob/main/0x07_Structs_and_Maps/01_Structs.go)

In this example we go over how structs can be utilized to hold different types of data. To start things off you may have noticed that the structure declaration is outside of the main function. Structures are declared outside of the main function for an important reason. It allows for methods. As you will see in the last example methods are a very useful feature to have when working with structs.

When creating the struct instance new_car, we see that the variables in the brackets are just values. This is because those values are assumed to represent the data to be stored as a struct instance. They are in the same order as declared in the struct and putting the wrong data in the wrong spot will result in either errors or incorrect data being stored.

The following example outlines how values can be added with a little more "accuracy".

## Adding_Values [</>](https://github.com/Syssos/Learning_Go/blob/main/0x07_Structs_and_Maps/02_Adding_Values.go)

When creating an instance of a struct, having data to fill the attributes is not required. This is because the compiler will automatically assign a "None" value for that type, i.e. a string will be empty "", an int will be 0.

We can add these values based off of their name using dot notation. By using the instance name dot attribute we can set the value stored there. This will be a useful way of storing, retrieving and manipulating data stored within the structs.

## Removing_Values [</>](https://github.com/Syssos/Learning_Go/blob/main/0x07_Structs_and_Maps/03_Removing_Values.go)

Before you start saying that we aren't removing values by setting them to 0, I just want to say, that's exactly what's happening in this example. The reason for that is because we cannot actually remove the "Ratings” attribute from the struct, without saving the data into another struct that doesn't contain it.

This is because at runtime the structs type definition cannot be changed. This means there is no way for the code to "update" the values a struct holds on the fly. If a struct is not supposed to contain the data the only way to essentially remove it is by setting it to "empty".

## Conclusion

Structs are a very useful feature in Go. It may not be the most used, but it will definitely make your life easier when needing to store data of different types together. If they seem complex, remember, it is just an accumulation of what we have learned up to this point. The data it is saving is no different than the normal types, it follows the same rules, it is just in a place that can be combined with other data.
