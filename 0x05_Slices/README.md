# Slices

## Overview

Slices in Go are more common to use then array's as slices have no set length. Although they work in a very similar fasion arrays and slices work in different manors. The Go compiler will take an array name, length, and an optional capacity when creating the array, and then return a slice for us to work with.

Sense a slice is easier to work with we will use them more in the code we write. This directory will cover using a slice, adding, and removing items, as well as accessing specific locations within a slice. The examples above will be critical to understanding what the description is talking about, so be sure to take a look at the example before/or after reading the description to get a clear picture of whats being discussed.

## Slices [</>](https://github.com/Syssos/Learning_Go/blob/main/0x05_Slices/00_Slices.go)
As explained above slices work in similar ways to an array. This example will cover declaring and printing a slice.

We can see that when declaring the slice we use the array arr, that is because the slice we are creating will use that array. If we created the slice without using the parrent array we would have to manually reassign the values into it, as if we were creating an array.

Go has a nifty built in short hand method of telling the compiler what items from the array to use when making the slice. In the example we start creating the slice in a similiar mannor as an array, however we set it equal to the parent array with some numbers in a bracket next to it ``` arr[0:6] ```. Those numbers are refrencing which indexes from the parent array, to copy over.

## Start [</>](https://github.com/Syssos/Learning_Go/blob/main/0x05_Slices/01_Start.go)

In this example we cover the start of a slice. Every slice or array will start at index 0. This makes it a pretty easy task to get the first value of a slice or array.

The code in this example tells the story for us. We see multiple slices have been created and the first value can always be accessed with 0

## End [</>](https://github.com/Syssos/Learning_Go/blob/main/0x05_Slices/02_end.go)

This example shows that working with the end of the slice is a little more complicated as the number isn't always a set. That raises the question as to how to work around that. The answer is the built in len() function. The length function will return an integar value of the number of items in the slice.

There is an intresting thing to note thats happening in this example however. When we refrence the last number of the slice while printing we have the value returned by len minus 1. Why minus one from the length?

We minus the one because slices start at 0. A slice containing ints, 

[]ints{8, 3, 42} 

this slice will have a len of 3, because theres 3 items. However the last value in the slice is stored at index 2, ie.. 

sli[2] = 42

If we put the len of the slice in sli, sli[3], we would get an out of range error because that value hiigher then the last stored value.

## Append [</>](https://github.com/Syssos/Learning_Go/blob/main/0x05_Slices/03_Append.go)

The benifit of working with slices is we can easily change the amount of values it holds by appending values. If we are working with set size arrays, we cannot say "hey add this value". The comiler will give us an out of range error much like printing the slice at index slice length.

To append an item to a slice we use the append function. The first argument taken is the slice to append to, and the second is the value to append. Go does a pretty good job at getting that value in the slice if its of the same type. Thankfully Go does the leg work on this one.

## Remove [</>](https://github.com/Syssos/Learning_Go/blob/main/0x05_Slices/04_Remove.go)

This example highlights removing a value from a slice, and some good practices to keep in mind. Due to the way slices and array's are stored in memory they will most likely be clumped together, added one right after the other, resulting in an ordered set of values. When removing a value we dont want to have an empty spot in memory, seporating the values.

To prevent that we start by setting the value at the end of the slice, to the position in the slice we are removing. This will ensure no gaps in memory can be created.

Once we have saved the last value in the slice elsewhere in the slice, we can remove it from the slice. However we want to do one more thing before that, we want to clear the value it is holding. Once that value is removed we can replace that slice, with a slice one index shorter.

So why clear that last slice value? We want to ensure no memory leaks are possible. While in this case nothing serious would happen even if there was, its important to keep in mind and practice. While memory leaks will mainly happen when working with a slice that is permanent and the element temporary, the practice never hurts.

## Conclusion

Slices are much easier to work with then arrays as they are not set in size. This aspect makes them easier to pass to functions, easier to append, and remove from. There are many different ways to manipulate slices and array's and it's amazing practice to play around with the, and see what they can do. Slices can be used all over, even more so when using Go for web based projects.

04_Remove touch on memory leaks, if you are unsure what they are and would like to know what they are, you can start your research at the wikipedia page.

[Memory Leak Wiki](https://en.wikipedia.org/wiki/Memory_leak)

[Back To Top](#slices)