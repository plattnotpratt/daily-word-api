package main

import (
  "encoding/json"
  "net/http"
  "github.com/go-chi/chi/v5"
  "log"
)

type WordResponse struct {
  Word string `json:"word"`
}

func main() {
loadWordsFromFile("words.txt")

  r := chi.NewRouter()

  r.Get("/daily-word", func(w http.ResponseWriter, r *http.Request){
    word := getDailyWord()
    resp := WordResponse{Word : word}
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(resp)
    if err != nil{
      log.Fatal(err)
    }
  })

  r.Get("/random-word", func(w http.ResponseWriter, r *http.Request){
    word := getRandomWord()
    resp := WordResponse{Word : word}
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(resp)
    if err != nil{
      log.Fatal(err)
    }
  })

  log.Println("Server Started on :8080")
  err := http.ListenAndServe(":8080", r)

  if err != nil{
    log.Fatal(err)
  }
}
