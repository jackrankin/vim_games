/*

this is the most overcomplicated thing ever but its just a homegrown, grass fed, concurrent trie

this version of the trie + parser is slower than the no-go routine version

this one averages about 600ms on my 2019 mbp 13 inch (8gb ram)

the other averages around 450ms

*/


package main

import (
    "fmt"
    "bufio"
    "os"
    "time"
    "runtime"
    "sync"
)

var mutex sync.Mutex

type trieNodeV2 struct {
    children []trieNodeV2
    end bool 
}

func trieAppendV2(root *trieNodeV2, word string) {
    for idx := 0; idx < len(word); idx++ {
        root = &root.children[word[idx] - 'A']
        if root.children == nil {
            root.children = make([]trieNodeV2, 26)
        }
    }
    root.end = true
}

func validateWordV2(root *trieNodeV2, word string) bool {
    for idx := 0; idx < len(word); idx++ {
        if root.children == nil {
            return false;
        }
        root = &root.children[word[idx] - 'A']
    }
    return root.end
}


func workerV2(jobs <- chan string, wg *sync.WaitGroup, root *trieNodeV2) {
    defer wg.Done()
    for liner := range jobs {
        mutex.Lock()
        trieAppendV2(root, liner);
        mutex.Unlock()
    }
}

func makeTrieV2() *trieNodeV2 {

    root := trieNodeV2{
        children: make([]trieNodeV2, 26),
        end:      false,
    }
    
    file, err := os.Open("./words.txt")

    if err != nil {
        panic(err)
    }

    defer file.Close()
    
    startTime := time.Now() 
    
    runtime.GOMAXPROCS(2)
    jobs := make(chan string)
    var wg sync.WaitGroup

    for i := 0; i < 2; i++ {
        wg.Add(1)
        go workerV2(jobs, &wg, &root)

    }

    go func() {
        defer close(jobs)
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            jobs <- scanner.Text()
        }

    }()

    wg.Wait()

    fmt.Println("Successful parsing in", time.Now().Sub(startTime))
    return &root

}
