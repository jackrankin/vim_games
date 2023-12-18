// this is the most overcomplicated thing ever but its just a homegrown, grass fed, free range, DIY trie

package main

import (
    "fmt"
    "bufio"
    "os"
)

type trieNode struct {
    letter byte
    children []trieNode
    end bool 
}

func trieAppend(root *trieNode, word string) {
    for idx := 0; idx < len(word); idx++ {
        root.children[word[idx] - 'A'].letter = word[idx]
        root = &root.children[word[idx] - 'A']
        if len(root.children) == 0 {
            root.children = make([]trieNode, 26)
        }

    }
    root.end = true
}

func validateWord(root *trieNode, word string) bool {
    for idx := 0; idx < len(word); idx++ {
        if len(root.children) == 0 || root.letter == 0 {
            return false;
        }
        root.children[word[idx] - 'A'].letter = word[idx]
        root = &root.children[word[idx] - 'A']
    }
    return root.end
}

func makeTrie() *trieNode {

    root := trieNode{
        letter:   '1',
        children: make([]trieNode, 26),
        end:      false,
    }

    file, err := os.Open("./words.txt")

    if err != nil {
        panic(err)
    }

    defer file.Close()
    
    scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text() 
        trieAppend(&root, string(line))
	}
    fmt.Println("Successful Parsing")
    return &root

}
