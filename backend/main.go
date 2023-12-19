package main

import (
    "net/http"
    "fmt"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/rs/cors"
)

var root = makeTrie();

func main() {
    server()
}

func server(){
    r := chi.NewRouter()
	r.Use(middleware.Logger)
    c := cors.Default()

	r.Get("/generateRandom", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(generateBoard()))
	})

	r.Get("/validateWord/{word}", func(w http.ResponseWriter, r *http.Request) {
        word := chi.URLParam(r, "word")
        fmt.Println(word)
        if validateWord(root, word) {
            w.Write([]byte("true"))
        } else {
            w.Write([]byte("false"))
        }
	})
    
    handler := c.Handler(r)

    http.ListenAndServe(":8000", handler)
}


