package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "log"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/rs/cors"
)


type Game struct {
    GameId string `json: "gameId"`
    GameString string  `json: "gameString"`
}

func serveGenerateBoard (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(generateBoard()))
}

func serveValidateWord(w http.ResponseWriter, r *http.Request) {
    word := chi.URLParam(r, "word")
    fmt.Println(word)
    if validateWord(&head, word) {
        w.Write([]byte("true"))
    } else {
        w.Write([]byte("false"))
    }
}

func makeRoom(w http.ResponseWriter, r *http.Request) {
    gameString := generateBoard()
    gameId := createGame(gameString)

    g := Game{
        GameId: gameId,
        GameString: gameString,
    }

    jsonData, err := json.Marshal(g)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(jsonData))
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}

func joinRoom(w http.ResponseWriter, r *http.Request) {
    name := chi.URLParam(r, "name")
    gameId := chi.URLParam(r, "gameId")

    gameString := joinGame(gameId)
    initUser(name, gameId)

    w.Write([]byte(gameString))
}


func finishGame(w http.ResponseWriter, r *http.Request){ 
    name := chi.URLParam(r, "name")
    gameId := chi.URLParam(r, "gameId")
    score :=  (chi.URLParam(r, "score"))

    addGame(name, gameId, score)

    leaderboard := getLeaderboard(gameId)

    jsonLB, err := json.Marshal(leaderboard)
    check(err)

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonLB)
}

func checkFinish(w http.ResponseWriter, r *http.Request){
    name := chi.URLParam(r, "name")
    gameId := chi.URLParam(r, "gameId")
    
    finish := checkUserFinish(name, gameId)
    if (finish == 1) {
        w.Write([]byte("true"))
    } else {
        w.Write([]byte("false"))
    }
}



func runServer(){

    r := chi.NewRouter()
    c := cors.Default()

	r.Use(middleware.Logger)
	r.Get("/generateRandom", serveGenerateBoard)
	r.Get("/validateWord/{word}", serveValidateWord)
	r.Get("/makeRoom", makeRoom)
	r.Get("/joinRoom/{name}/{gameId}", joinRoom)
    r.Get("/finish/{name}/{gameId}/{score}", finishGame)
    r.Get("/checkFinish/{name}/{gameId}", checkFinish)

    handler := c.Handler(r)

    http.ListenAndServe(":8000", handler)
}

