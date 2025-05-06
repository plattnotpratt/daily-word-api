package main

import (
  "crypto/sha1"
  "math/rand/v2"
  "encoding/hex"
  "time"
  "bufio"
  "log"
  "os"
)

var words []string

func loadWordsFromFile(path string){
  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    word := scanner.Text()
    if word != ""{
      words = append(words, word)
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

func getRandomWord() string{
  index := rand.IntN(len(words))
  return words[index]
}

func getMultipleWords(x int) []string{
  var words [0]string
  for i := 0; i < x; i++{
    words.append(getRandomWord())
  }
  return words
}
func getDailyWord() string{
  date := time.Now().Format("2006-01-02")
  h := sha1.New()
  h.Write([]byte(date))
  hash := hex.EncodeToString(h.Sum(nil))

  index := int(hash[0]) % len(words)
  return words[index]
}
