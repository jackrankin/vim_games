package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var db *sql.DB

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func initUser(name string, gameId string) {
    fmt.Println("INSERTING NAME:", name)
    _, err := db.Exec("INSERT INTO users (game_id, username, finish_time) VALUES ($1, $2, $3)", gameId, name, 1e9)
    check(err)
}

func createGame(gameString string) string {
    fmt.Println("CREATING GAME:", gameString)
    _, err := db.Exec("INSERT INTO games (game_string) VALUES ($1)", gameString)
    check(err)

    rows, err := db.Query("SELECT game_id FROM games WHERE game_string = $1", gameString)
    check(err)

    var gameId string
    for rows.Next() {
        err := rows.Scan(&gameId)
        check(err)
    }
    return gameId
}

func joinGame(gameId string) string {
    fmt.Println("JOINING GAME:", gameId)
    rows, err := db.Query("SELECT game_string FROM games WHERE game_id = $1", gameId)
    check(err)
    var gameString string
    for rows.Next() {
        err := rows.Scan(&gameString)
        check(err)
    }
    return gameString
}

func finishGame(name string, time float32) {
    _, err := db.Exec("UPDATE users SET finish_time = $1 WHERE username = $2", time, name)
    check(err)
}

func connectDatabase() {
    var err error
    err = godotenv.Load("../.env")
    check(err)

    connStr := os.Getenv("DB_CONNECTION_STRING")

    db, err = sql.Open("postgres", connStr)
    check(err)

}
