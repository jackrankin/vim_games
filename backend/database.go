package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

type User struct {
    Username string `json: "username"`
    Score string `json: "score"`
}

var db *sql.DB

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func initDB() {
    _, err := db.Exec("DROP TABLE IF EXISTS users; DROP TABLE IF EXISTS games; DROP SEQUENCE IF EXISTS sequence_thousand CASCADE;")
    check(err)
    _, err = db.Exec("CREATE TABLE games (game_id SERIAL PRIMARY KEY, game_string VARCHAR(16) NOT NULL);")
    check(err)
    _, err = db.Exec("CREATE TABLE users (user_id SERIAL PRIMARY KEY,game_id INT REFERENCES games(game_id) ON DELETE CASCADE,username VARCHAR(255) NOT NULL,score INT,finished INT DEFAULT 0);")
    check(err)
    _, err = db.Exec("CREATE SEQUENCE sequence_thousand START WITH 2000;")
    check(err)
    _, err = db.Exec("ALTER TABLE games ALTER COLUMN game_id SET DEFAULT nextval('sequence_thousand');")
    check(err)
}

func initUser(name string, gameId string) {
    fmt.Println("INSERTING NAME:", name)
    _, err := db.Exec("INSERT INTO users (game_id, username, score) VALUES ($1, $2, $3)", gameId, name, 1e9)
    check(err)
}

func checkUserFinish(name string, gameId string) int {
    fmt.Println("CHECKING IF USER PLAYED:", name)
    rows, err := db.Query("SELECT finished FROM users WHERE game_id = $1 AND username = $2", gameId, name)
    check(err)

    var finish int
    for rows.Next() {
        err := rows.Scan(&finish)
        check(err)
    }
    return finish

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

func addGame(name string, gameId string, time string) {
    _, err := db.Exec("UPDATE users SET score = $1, finished = 1 WHERE username = $2 AND game_id = $3", time, name, gameId)
    check(err)
}

func getLeaderboard(gameId string) []User {
    rows, err := db.Query("SELECT users.username, users.score FROM users JOIN games ON users.game_id = games.game_id WHERE games.game_id = $1 ORDER BY users.score DESC", gameId)
    check(err)
    var username string
    var score string 
    leaderboard := make([]User, 0)
    for rows.Next() {
        err := rows.Scan(&username, &score)
        check(err)
        u := User{
            Username: username,
            Score: score,
        }
        leaderboard = append(leaderboard, u)
    }
    return leaderboard;
}

func connectDatabase() {
    var err error
    err = godotenv.Load("../.env")
    check(err)

    connStr := os.Getenv("DB_CONNECTION_STRING") 

    db, err = sql.Open("postgres", connStr) // I'm pretty sure this will work with any sql database if u swap the postgres string
    check(err)

    initDB()

}
