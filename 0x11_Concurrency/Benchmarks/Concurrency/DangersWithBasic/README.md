# Dangers with Basic Concurrency usecases

In [example.go](https://github.com/Syssos/Learning_Go/blob/main/0x11_Concurrency/Benchmarks/Concurrency/DangersWithBasic/example.go) there are 2 functions that are going to be tested. Both of them with different approach's to handling the same task, finding a name that starts with a specific letter in a list.

## About Tests

In the function <b>FindUserStartingWith()</b> everything is handled synchronously and has no concurrency. There is some minor string manipulation done on each string, to help assist in matching no matter the case, and to compare the first letter of the string. 

In the <b>FindUserStartingWithConcurrent()</b> function, the task is handled with concurrency in mind. The dataset is split in half, and there are 2 goroutines that handle looking for the name. When either of the go routines finds a name, it sends the value into the chanel. With the help of a select statement the first result to be passed into the channel is returned, after sending a quit signal to the possibly remaining goroutine via another channel.

Each function will be tested with 2 sets of name lists. The first list tested against will be relativaly shory only containing a handful of names. The second list will contain over 18000 names.

There is little difference in the actual logic that handles the searching, and the lists will be the exact same. Which means the difference in the results is going to be coming from the method in which the logic is applied.

## Results

```
go test -bench=.
```

```
goos: linux
goarch: amd64
pkg: github.com/Syssos/Advanced_Go/Benchmarks/Concurrency/DangersWithBasic
cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
BenchmarkFindUserStartingWith-4                         19463281                59.06 ns/op
BenchmarkFindUserStartingWithConcurrent-4                 763920              1726 ns/op
BenchmarkFindUserStartingWithLong-4                       785029              1525 ns/op
BenchmarkFindUserStartingWithConcurrentLong-4             162685              6956 ns/op
PASS
ok      github.com/Syssos/Advanced_Go/Benchmarks/Concurrency/DangersWithBasic   4.990s

```

The results above do not appear to be in Concurrencies favor. There is however a very valid explination for these results.

## Result Discussion

The reason for the results seen, has to do with all of the overhead that comes with creating goroutines. Golang's benchmark testing will itterate a loop b.N amount of times. What this means is the efficiency of our code goes beyond a one time use.

While in theory the concurrent approach appears faster by breaking the data up and checking each section in a goroutine, by the time both goroutines and there resources are ready to start looking, the synchronous method could have finished already. In this case that is whats happening.

> Something further to note is the larger the list of names got, the time for each op got much larger in comparison. The synchronous approach got roughly <b>25x</b> slower when dealing with the long list. Opposed to the concurrent approach only getting <b>4x</b> slower.

## So why use Goroutines?

While the results from this benchmark may not show it, goroutines can deffinetly help speed up normally slow approaches. The [Advanced Example Benchmark](https://github.com/Syssos/Learning_Go/tree/main/0x11_Concurrency/Benchmarks/Concurrency/AdvancedExampleBenchmark) is a great display of goroutines, when used properly, shortening the time it takes to complete a task. That example is unlike this one in the sense that it doesn't suffer from the same flaw.

This example's flaw, is less of a flaw and more of an approach issue. It is better left to as a synchronous approach.

While splitting up the dataset could improve the chance of finding the result faster, we are only after the first result. This means that even if we gain the time benifit of finding the result in the second half of the split list, we are still making up for the time lost generating both the goroutines and allocating resources. This example was not intended to show results of goroutines saving time, but results that display the effects of using goroutines poorly, or when not needed.

