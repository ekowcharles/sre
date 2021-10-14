package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func buildResponse(c int) string {
	r := Payload{strconv.Itoa(c), http.StatusText(c)}
	JSON, _ := json.Marshal(r)

	return string(JSON)
}

func main() {
	defer fmt.Println("Exited!")

	port := getEnv("PORT", "8040")
	version := App{getEnv("APP_VERSION", "0.0.0")}

	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		JSON, _ := json.Marshal(version)

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(JSON))
	})

	http.HandleFunc("/random", func(w http.ResponseWriter, rq *http.Request) {
		i := rand.Intn(sl)
		s := ss[i]

		w.WriteHeader(s)
		io.WriteString(w, buildResponse(s))
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
