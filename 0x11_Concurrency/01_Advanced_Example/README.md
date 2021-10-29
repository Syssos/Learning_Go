# Advanced Example of Concurrency [</>](https://github.com/Syssos/Learning_Go/blob/main/0x11_Concurrency/01_Advanced_Example/example.go)

This example is to outline a perfect time for a developer to take advantage of Goroutines. Within [example.go](https://github.com/Syssos/Learning_Go/blob/main/0x11_Concurrency/01_Advanced_Example/example.go) there are 3 calls to API's. The response from each request is then appended to a list, and returned when all 3 requests are done.

## About Example 

This test is designed to show 2 things. The difference in complexity/readbility when incorperating concurrent approaches, and the difference in speed gained from concurrent approaches.

There are 2 functions to focus attention on. The first of which is <b>FetchThreeJokes()</b>. This function will <b>synchronously</b> send 3 HTTP requests and return the data for each in one list. This method forces each request to be made in order, one after the other.

In the function <b>FetchThreeJokesConcurrent()</b> the same work is being done, however in a concurrent mannor. This means that each API request is handled by a goroutine, resulting in each request being scheduled to run whenever there is downtime in other goroutines.

This scheduling is what allows for the concurrent approach to be upwards of <b>3x</b> faster.

## Running the Example

```
go run ./example.go
```

```
A guy walks into a bar and asks for 1.4 root beers.
The bartender says "I'll have to charge you extra, that's a root beer float".
The guy says "In that case, better make it a double."
I bought some shoes from a drug dealer. I don't know what he laced them with, but I was tripping all day!
What does the MacBook have in common with Donald Trump?

I would tell you....
But I don't compare apples to oranges.
My wife and I have reached the difficult decision that we do not want children.
If anybody does, please just send me your contact details and we can drop them off tomorrow.
Oysters hate to give away their pearls because they are shellfish.
Judge: "I sentence you to the maximum punishment..."
Me (thinking): "Please be death, please be death..."
Judge: "Learn Java!"
Me: "Damn.
```

Above is an example of the output a user may see when running the example. There should be 6 jokes in total. Two for Programming, two for Puns, two for Misc.

## Benchmarking

This example has its own benchmark test in the [Benchmarks](https://github.com/Syssos/Learning_Go/tree/main/0x11_Concurrency/Benchmarks) directory, that show off the benifit of this approach. If you are intrested in the tests, or the results you can check out [this](https://github.com/Syssos/Learning_Go/tree/main/0x11_Concurrency/Benchmarks/Concurrency/AdvancedExampleBenchmark) directory.