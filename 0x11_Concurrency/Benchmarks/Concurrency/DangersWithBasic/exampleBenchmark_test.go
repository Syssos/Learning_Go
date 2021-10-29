package main

import (
    // "fmt"
    "testing"
    "encoding/json"
    "io/ioutil"
)

var Names NamesList
var Results string

func init() {
    file, _ := ioutil.ReadFile("names.json")

    var data NamesList
    _ = json.Unmarshal([]byte(file), &data)

    Names = data
}

func BenchmarkFindUserStartingWith(b *testing.B) {
    var r string
    for i := 0; i < b.N; i++ {
        r = FindUserStartingWith("h", Names.Short)
    }

    Results = r
}

func BenchmarkFindUserStartingWithConcurrent(b *testing.B) {
    var r string
    for i := 0; i < b.N; i++ {
       r = FindUserStartingWithConcurrent("h", Names.Short)
    }

    Results = r
}

func BenchmarkFindUserStartingWithLong(b *testing.B) {
    var r string
    for i := 0; i < b.N; i++ {
        r = FindUserStartingWith("h", Names.List)
    }

    Results = r
}

func BenchmarkFindUserStartingWithConcurrentLong(b *testing.B) {
    var r string
    for i := 0; i < b.N; i++ {
        r = FindUserStartingWithConcurrent("h", Names.List)
    }

    Results = r
}