package main

import (
    "fmt"
    "math/rand"
)

func generateBoard() string {
    board := ""

    var blocks [][]string
    blocks = append(blocks, []string {"R", "I", "F", "O", "B", "X"})
    blocks = append(blocks, []string {"I", "F", "E", "H", "E", "Y"})
    blocks = append(blocks, []string {"D", "E", "N", "O", "W", "S"})
    blocks = append(blocks, []string {"U", "T", "O", "K", "N", "D"})
    blocks = append(blocks, []string {"H", "M", "S", "R", "A", "O"})
    blocks = append(blocks, []string {"L", "U", "P", "E", "T", "S"})
    blocks = append(blocks, []string {"A", "C", "I", "T", "O", "A"})
    blocks = append(blocks, []string {"Y", "L", "G", "K", "U", "E"})
    blocks = append(blocks, []string {"A", "B", "M", "J", "O", "A"})
    blocks = append(blocks, []string {"E", "H", "I", "S", "P", "N"})
    blocks = append(blocks, []string {"V", "E", "T", "I", "G", "N"})
    blocks = append(blocks, []string {"B", "A", "L", "I", "Y", "T"})
    blocks = append(blocks, []string {"E", "Z", "A", "V", "N", "D"})
    blocks = append(blocks, []string {"R", "A", "L", "E", "S", "C"})
    blocks = append(blocks, []string {"U", "W", "I", "L", "R", "G"})
    blocks = append(blocks, []string {"P", "A", "C", "E", "M", "D"})

    for i := 0; i < 16; i++ {
        board += blocks[i][rand.Intn(6)]
    }

    fmt.Println(board)

    return board
}




