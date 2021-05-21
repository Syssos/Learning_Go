# Methods and Interfaces

## Overview

Methods in Go are basically short hand ways of running a function designed to work with a specific struct. In the examples that follow well cover why and how methods are benificial for us as developers.

## 00_Methods

This example shows how methods are created and used. While the example is just printing the name of the "car" in a preset string, it demenstrates what a method would be used for.

Methods are designed to easily use the same function with any instance of a struct. This example shows that both "new_car", and "another_one" could use the function "print_name" via dot notation.

Another benifit to working with methods is they have access to all the contents inside of the struct. 

When calling a method, we use ```struct_instance.function()```.

When declaring the method we use ``` func (instance_name struct_type) function_name() return_type```, where "instance_name" is a variable cantaining the same data as "struct_instance"

We can use "instance_name" inside of the method function to retreive a variable specific to the struct that called it. We see this in the example as each instance of car has a different output, when the method is called. 

## Interfaces

Interfaces are a handy way to name a collection of methods. To help make this concept make a little more sense, lets look at this in a different more relatable manor then the example shows. 

First lets pretend were making a video game, and we know we want there to be bad guys and good guys.

In this game every good guy and bad guy takes damage, say via ``` char.damage(21) ```. We also need to update the screen so the player knows the health values, lets say ``` char.updateDisplay() ``` handles that.

In any case of either a good, or bad guy taking damage each of these functions need to run. We can use an interface to make this task easier. An example for this case would be ``` damageChar(char) ```.

In the example we are using geometry as a way to demenstrate interfaces. The methods are no different then the previous example, which is a nice touch because we don't have to account for an interface when creating the methods.

When creating an interface we declare it in a simular manor to a struct. We start with the keyword type, then the interface name, followed by the interface keyword. Inside of the brackets are the methods and the return values datatype for those methods, the interface will store.

The example outlines that we can create an instance of an interface via ```new_inst := interface_name(struct_name) ```, and use it to call the methods in a simular manor to how we would if we called the methods from the struct instance itself(``` struct_name.function() ```, ``` interface_name.function() ```). While this is good to know and can still be useful, the real power of this comes from using them in functions.

By Creating functions that take an argument of type "instance_name", you can pass multiple types of structs to that function so long they contain the corresponding methods supported by the interface. This can be an very convienent to shorten the amount of code needed to accomplish a task.

## Conclusion

While at first interfaces may be a hard concept to grasp they are amazing at saving time when it comes to working with multiple types of simular structs. This concept will take practice to give you the best chance at understanding how to utilize them. If you need more material [Go by Example](https://gobyexample.com/interfaces) has some great examples, and so does [A Tour of Go](https://tour.golang.org/methods/1).