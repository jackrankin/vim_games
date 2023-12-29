package main

func main() {
    connectDatabase()
    makeTrie()
    runServer()
    defer db.Close()
}
