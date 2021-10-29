# Using Channels with Goroutines

Using channels can become on of the most benifical ways of utilizing Goroutines when utilizing the correct approaches. Channels should be considered as pipes that connect various Go routines together.

When utilizing a Goroutine if the data is not sent via a channel, then the data will be copied, and a copy of the data will have the work done on it. This was intentional by the creators to prevent accidental data manipulation.

In the example you will see that there are 2 use cases of Goroutines, each doing work on seprate parts of the dataset, but communicating through the same channel to send results back.

While this example is perfect for us to see what is going on, note that it is not quite how we would want to use go routines, in the sense of redundancy due to the amount, and type of information.

A more complex version of this example could include a large name list, and an objective to find a name that starts with "h". With a function that checks if a name starts with "h", the same concepts can be applied, the name list can be broken up, and passed to multiple go routines. Then the data can be sent back to the caller, or elsewhere via the channel.