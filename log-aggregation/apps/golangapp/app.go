package main

import (
  "encoding/json"
  "fmt"
  "io"
  "log"
  "net/http"
  "os"
)

type Payload struct {
  Code string `json:"code"`
  Description string `json:"desciption"`
}

type App struct {
  Version string `json:"version"`
}

func main() {
  defer fmt.Println("Exited!")

  port := getEnv("PORT", "8993")
  version := App{ getEnv("APP_VERSION", "0.0.0") }

  http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
    JSON, _ := json.Marshal(version)

    w.WriteHeader(http.StatusOK)
    io.WriteString(w, string(JSON))
  })

  http.HandleFunc("/exception", func(w http.ResponseWriter, rq *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)

    panic("Doing this only for the logs")
  })

  fmt.Printf("Listening on port %s ...\n", port)

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }

    return value
}