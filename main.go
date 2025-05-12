package main

import (
  "encoding/json"
  "net/http"
  "github.com/go-chi/chi/v5"
  "log"
  "strconv"
)

type WordResponse struct {
  Word string `json:"word"`
}

type WordsResponse struct {
  Words []string `json:"words"`
}

type MessageResponse struct {
  Message string `json:"message"`
}

func main() {
loadWordsFromFile("words.txt")

  r := chi.NewRouter()

  r.Get("/", func(w http.ResponseWriter, r *http.Request){
    resp := MessageResponse{Message : "This is a daily/random word api. For more information check out the github repo: https://github.com/plattnotpratt/daily-word-api"}
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(resp)
    if err != nil{
      log.Fatal(err)
    }
  })

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

  r.Get("/random-word/{length}", func(w http.ResponseWriter, r *http.Request){
    length, _ := strconv.Atoi(chi.URLParam(r, "length"))
    word := getWordOfLength(length)
    resp := WordResponse{Word : word}
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(resp)
    if err != nil{
      log.Fatal(err)
    }
  })

  r.Get("/random-words/{count}", func(w http.ResponseWriter, r *http.Request){
    count, _ := strconv.Atoi(chi.URLParam(r, "count"))
    words := getMultipleWords(count)
    resp := WordsResponse{Words : words}
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(resp)
    if err != nil{
      log.Fatal(err)
    }
  })

  r.Get("/random-words/{count}/{length}", func(w http.ResponseWriter, r *http.Request){
    count, _ := strconv.Atoi(chi.URLParam(r, "count"))
    length, _ := strconv.Atoi(chi.URLParam(r, "length"))
    words := getWordsWithLength(count, length)
    resp := WordsResponse{Words : words}
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
