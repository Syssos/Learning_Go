# Concurrency

> "In computer science, concurrency is the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or at the same time simultaneously partial order, without affecting the final outcome." - Wikipedia

With many different definitions of the word, Concurrency, writing concurrent code can appear to be a daunting task for first timers. However with a proper understanding of what it means to write concurrent code, it can actually be a very fun and rewarding experiance.

To start off lets look at what concurrency means in the feild of computer science, and how we can use this methodolgy to our advantage as programmers.

## Table of Contents

 - [Overview](#concurrency)
 - [Table of Contents](#table-of-contents)
 - [Concurrency in Go](#concurrency-in-go)
 - [Concurrency VS Parallelism](#concurrency-vs-parallelism)
 - [Concurrency Benifits](#benifits-of-concurrency)
 - [Extra Sources](#sources)

## Concurrency in Go

Concurrency is more of a mindset, then anything, when it comes to how a developer should incorperate concurrency into their code. What this means is that writing concurrent code is more about the steps taken in completing a task. Lets look at an analogy of concurrency in real-life. 

Lets say you are a high school teacher, and you have 7 classes a day, each class 50 minutes, with a 10 minute period inbetween classes.

One day you get a notice about some tax related things and get stuck with an hour worth of paperwork to fill out. This paperwork must be submited to the tax office which closes 30 minutes after you get out of work, just enough time to get there. 

A non-concurrent approach would mean you have to sit down for exactly one hour to fill out the paperwork. In order for you to make it to the office in time to submit the paperwork you would have to cancel one class at the least.

A concurrent approach would incorperate be the ability to "Break up" the paperwork, lets say there is 6 pages to the paperwork, each requiring 10 minutes to complete. Doing one page, each 10 minute time period inbetween classes, would give you time to complete all the paperwork without having to cancel any classes. This works because the paperwork can be "schedualed" for when your students were doing "other things".

> Go will breaking "work" up over multiple light weight "threads" known as Goroutines, and schedule them to run whenever there is available resources. 

In the anology above there are 2 major "events" we are keeping track of, "Classes" and the "Paperwork". Programatically these "events" would need to be started in some way (ie, do_class(), or do_paperwork()). Note that the function that starts the work is whats desired to be sceduled as a Goroutine. This is different then say scheduling a portion of do_paperwork(), like do_page_1(), as it allows more of the processing to be incorperated in the scope of the Goroutine.

![image for go](https://raw.githubusercontent.com/Syssos/Learning_Go/main/0x11_Concurrency/etc/images/ConcurrencyInGo.png)

With the Goroutines configured in the mannor outlined above, we give the Go schedualer the ability to structure the work in a similar mannor to the teacher in the example. When there is not a class being taught, then the paperwork can be processed. This allows for us to do both of these tasks in a much more effiecent way, saving time in the end as there is no wasted time.

## Concurrency VS Parallelism

When first learning Go's concurrency model, it can often times be confused with parallelism. This is because the topics are simular to an extent, but differ in some specific ways. Parallelism is the act of processing information all at the same time. Think of a multi-core processor, it can handle multiple tasks at once.

![image for multi core processing](https://raw.githubusercontent.com/Syssos/Learning_Go/main/0x11_Concurrency/etc/images/parallelism.jpg)

Concurrency is the act of scheduling tasks to run as need be, or when there is an availability in processing power, this can be compared to a single core processor, that must break various "work" up over one core.

![image for single core processing](https://raw.githubusercontent.com/Syssos/Learning_Go/main/0x11_Concurrency/etc/images/concurrency.jpg)


While they both can have some "work" being done at the "same time" based on modern computers, the overall concept is different. This is an important distinction to note.

## Benifits of Concurrency

Concurrency is something that can be greatly benifited from when used in the right context. It can allow for one to optimize their code by eliminating any time wasted in the processing processes.

While there are percs to goroutines there can also be drawbacks if goroutines are used poorly. The example within the benchmark test folder [Dangers with Basic](https://github.com/Syssos/Learning_Go/tree/main/0x11_Concurrency/Benchmarks/Concurrency/DangersWithBasic) shows that when a concurrent approach is attempted, when not needed, there can actually be a harmful impact on the performance of the code.

However when used properly, there can be huge gains in performance, as seen in the benchmark test for [Advanced Example](https://github.com/Syssos/Learning_Go/tree/main/0x11_Concurrency/Benchmarks/AdvancedExampleBenchmark).

## Sources

- [Sync Waitgroup Example](https://resources.oreilly.com/examples/9781783983483/blob/master/3483OS_Code/3483OS_01_Codes/ch1_1_patientgoroutine.go) <!-- This is really nice code ngl -->

- [Images for concurrency & parallelism](https://medium.com/rungo/achieving-concurrency-in-go-3f84cbf870ca) <!-- Found on Google -->

[Back to Top](#concurrency)