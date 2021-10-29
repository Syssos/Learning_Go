package main

import (
    "fmt"
    "sync"
    "strings"
    "encoding/json"
    "io/ioutil"
)

// Used to pull names from file "names.json"
type NamesList struct {
    List []string `json: "list"`
    Short []string `json: "short"`
}

// Non Concurrent approach to finding users
func FindUserStartingWith(c string, dataset []string) (result string) {
    check := strings.ToLower(c)
    for _, item := range(dataset){
        lowered := strings.ToLower(item)
        if strings.HasPrefix(lowered, check) {
            return item
        }
    }
    return ""
}

// Concurrent approach to finding users
func FindUserStartingWithConcurrent(c string, dataset []string) (result string) {
    res := make(chan string, 2)
    qt := make(chan int, 2)

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        if result := NameFindingWorker(c, dataset[:len(dataset)/2], qt); result != "" {
            res <- result
        }
        wg.Done()
    }()
    go func() {
        if result := NameFindingWorker(c, dataset[len(dataset)/2:], qt); result != "" {
            res <- result
        }
        wg.Done()
    }()

    select {
    case data := <-res:
        qt <- 0
        return data
    }

    wg.Wait()
    return ""
}

// Helper function called by multiple go routines.
func NameFindingWorker(c string, dataset []string, qt chan int) string {
    check := strings.ToLower(c)
    for _, item := range(dataset){
        select {
        case <- qt:
            return ""
        default:
            lowered := strings.ToLower(item)
            if strings.HasPrefix(lowered, check) {
                return item
            }
        }
    }

    return ""
}

func main() {
    file, _ := ioutil.ReadFile("./names.json")

    var data NamesList
    _ = json.Unmarshal([]byte(file), &data)

    nameSN := FindUserStartingWith("h", data.Short)
    nameLN := FindUserStartingWith("h", data.List)
    nameSC := FindUserStartingWithConcurrent("h", data.Short)
    nameLC := FindUserStartingWithConcurrent("h", data.List)

    fmt.Println("Result for finding name non-concurrently (Short list): ", nameSN)
    fmt.Println("Result for finding name non-concurrently (Long list): ", nameLN)
    fmt.Println("Result for finding name concurrently (Short list): ", nameSC)
    fmt.Println("Result for finding name concurrently (Long list): ", nameLC)
}