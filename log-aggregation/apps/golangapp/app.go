package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func buildResponse(w http.ResponseWriter, c int) {
	r := Payload{strconv.Itoa(c), http.StatusText(c)}
	JSON, _ := json.Marshal(r)

	w.WriteHeader(c)
	io.WriteString(w, string(JSON))
}

func main() {
	defer logln(Debug, "Exited!")

	port := getEnv("PORT", "8040")
	version := App{getEnv("APP_VERSION", "0.0.0")}

	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		JSON, _ := json.Marshal(version)

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(JSON))
	})

	http.HandleFunc("/http/", func(w http.ResponseWriter, rq *http.Request) {
		exp := regexp.MustCompile(`/http/(?P<id>\b[0-9]{3}$)`)
		match := exp.FindStringSubmatch(rq.URL.Path)

		if len(match) != 2 {
			buildResponse(w, http.StatusBadRequest)

			return
		}

		c, err := strconv.Atoi(match[1])
		if err != nil || http.StatusText(c) == "" {
			buildResponse(w, http.StatusBadRequest)

			return
		}

		buildResponse(w, c)
	})

	http.HandleFunc("/random", func(w http.ResponseWriter, rq *http.Request) {
		i := rand.Intn(sl)

		buildResponse(w, ss[i])
	})

	http.HandleFunc("/exception", func(w http.ResponseWriter, rq *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)

		panic("Doing this only for the logs")
	})

	logf(Debug, "Listening on port :%s", port)
	fmt.Print(&buf)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
