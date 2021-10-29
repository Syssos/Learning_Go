# Unit Tests

## Overview

Unit tests can be extremely an extremly helpful tool for ensuring code's functionality and performance. Due to testing code being common amongst programmers, the developers of Go have created a way to do that with the help of the "testing" package. While this package contains a majority of the tools needed for developers to test their code, it still leaves test cases, and validating output up to the developer. This repository is targeted at explaining how these tools can be used to create efficient ways of checking code.

## About the Testing Package

The testing package was introduced to the standard library to give developers the option to automate the testing process for their code. This package does all of the leg work, while leaving the actual test cases up to the developer allowing them to test a very large range, if not all, of the components in a package.

This package includes things such as functionality testings, benchmark testings, sub-tests/benchmarks, the ability to test the Main function, and a multitude of little, but benificial components.

Test files are meant to be part of the package as they can be a vital way to check that none of the codes functionality or performance was impacted when making changes. Each testing file delcared package should be the same as the package the functions desired to test are in.

## add [</>](https://github.com/Syssos/Learning_Go/blob/main/0x10_Unit_Tests/00_add.go)

This file contains the code that will be tested in against in the unit tests. Due to how the Golang testing library is configured, we do not have to do anything special to account for the tests when initial writing the code.

## add_test [</>](https://github.com/Syssos/Learning_Go/blob/main/0x10_Unit_Tests/01_add_test.go)

Test file names are denoted by the suffix "<b>_test</b>". This tells the Go compiler that there are test cases inside of the file to be used while testing. Inside this file you will notice 2 "types" of functions, one starting with the prefix "<b>Test</b>", the other with the prefix "<b>Benchmark</b>".

The testing package will automatically detect these functions based on this prefix. Each function starting with "Test" will be used to check functionality, whereas functions starting with "Benchmark" will be used to check for performance.

## Running Tests

### Running functionality tests
To run the test within this package, you can utilize Golangs cli test command.

```
$ go test .
ok      github.com/Syssos/unit_t        0.002s
```
Be sure to include the period to tell the go runtime to run all tests. The output should be a single line, letting you know that no test cases had failed.

A specific test can be told to run with the "-run" argument followed by the tests function name

```
$ go test -run TestSum
PASS
ok      github.com/Syssos/unit_t        0.002s
```
Unlike the last test, running a single test will let you know it passed by printing the phrase "PASS" above the results.

If you would like a more verbose output you can make use of the "-v" flag.

```
$ go test -v .
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestDif
--- PASS: TestDif (0.00s)
ok      github.com/Syssos/unit_t        0.002s
```
The output should have more data letting you know each functions testing time and its status.

### Running Benchmark Tests

Simular to the functionality tests you will use the "test" command, however in order to run the benchmark tests, you must specify the "-bench" flag, as well as which tests to run.

Using a period denotes all tests as ones that should run.

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/Syssos/unit_t
cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
BenchmarkSum-4          1000000000               0.3715 ns/op
BenchmarkDif-4          1000000000               0.3672 ns/op
PASS
ok      github.com/Syssos/unit_t        0.826s
```

If a specific function name is given, then only that test will be ran.

```
$ go test -bench=BenchmarkDif
goos: linux
goarch: amd64
pkg: github.com/Syssos/unit_t
cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
BenchmarkDif-4          1000000000               0.3812 ns/op
PASS
ok      github.com/Syssos/unit_t        0.428s
```
One thing to note is that before running benchmark tests you will run the functionality tests. This can be seen when using the "-v" flag while running benchmark tests.

```
$ go test -v -bench=.
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestDif
--- PASS: TestDif (0.00s)
goos: linux
goarch: amd64
pkg: github.com/Syssos/unit_t
cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
BenchmarkSum
BenchmarkSum-4          1000000000               0.3771 ns/op
BenchmarkDif
BenchmarkDif-4          1000000000               0.3721 ns/op
PASS
ok      github.com/Syssos/unit_t        0.843s
```
## Conclusion

With how easy the standard testing library makes it to test code, it is by far one of the most help, "out of the box" features of Go. It is highly recommended that further research is done into unit testing. The topic is very broad and with each function having different goals, data types, and results to deal with, it's important to have good starting points. To learn more about unit tests or the "testing" package I would check out the official [Golang](https://golang.org/pkg/testing/) site and read up on what they have to say about testing.
