# 0x00 Concurrency, 01 Advanced Example Benchmark

This is the benchmark testing for the advanced example of concurrency, outlining benifits to using concurrent approachs when approapriate.

```
goos: linux
goarch: amd64
pkg: Advanced_Go/Benchmarks/Concurrency/AdvancedExampleBenchmark
cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
BenchmarkFetchThreeJokes-4                     2         972044776 ns/op
BenchmarkFetchThreeJokesConcurrent-4           3         408747316 ns/op
PASS
ok      Advanced_Go/Benchmarks/Concurrency/AdvancedExampleBenchmark     5.537s

```