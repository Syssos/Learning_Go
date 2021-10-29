package main

import (
    "testing"
)

var Results []string

func BenchmarkFetchThreeJokes(b *testing.B) {
    var r []string
    for i := 0; i < b.N; i++ {
        r = FetchThreeJokes()
    }

    Results = r
}

func BenchmarkFetchThreeJokesConcurrent(b *testing.B) {
    var r []string
    for i := 0; i < b.N; i++ {
        r = FetchThreeJokesConcurrent()
    }

    Results = r
}