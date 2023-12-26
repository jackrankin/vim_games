package main

import (
    "fmt"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/rs/cors"
	"github.com/olahol/melody"
)

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

func runServer(){

    r := chi.NewRouter()
    m := melody.New()
    users := make(map[string]int)
    id := 0

	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("Client connected")
        s.Set("id", id)
        users[id] = s
        id++
	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("Client disconnected")
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Printf("Message from client: %s\n", msg)
		m.Broadcast(msg)
        m.Broadcast()
	})

    c := cors.Default()

	r.Use(middleware.Logger)
	r.Get("/generateRandom", serveGenerateBoard)
	r.Get("/validateWord/{word}", serveValidateWord)
    r.Get("/websocket", func (w http.ResponseWriter, r *http.Request){
        m.HandleRequest(w, r)
    })

    handler := c.Handler(r)

    http.ListenAndServe(":8000", handler)
}

