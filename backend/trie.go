package main

import (
    "fmt"
    "bufio"
    "os"
    "time"
)

var head trieNode;

type trieNode struct {
    children []trieNode
    end bool 
}

func trieAppend(root *trieNode, word string) {
    for idx := 0; idx < len(word); idx++ {
        root = &root.children[word[idx] - 'A']
        if root.children == nil {
            root.children = make([]trieNode, 26)
        }
    }
    root.end = true
}

func validateWord(root *trieNode, word string) bool {
    for idx := 0; idx < len(word); idx++ {
        if root.children == nil {
            return false;
        }
        root = &root.children[word[idx] - 'A']
    }
    return root.end
}

func makeTrie() {
    
    head = trieNode{
        children: make([]trieNode, 26),
        end:      false,
    }
    
    file, err := os.Open("./words.txt")

    if err != nil {
        panic(err)
    }

    defer file.Close()
    
    startTime := time.Now() 

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        trieAppend(&head, scanner.Text())
    }

    fmt.Println("Successful parsing in", time.Now().Sub(startTime))

}
