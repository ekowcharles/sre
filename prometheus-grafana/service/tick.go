package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	defer fmt.Println("Exited!")

	port := os.Getenv("PORT")
	app := os.Getenv("APPLICATION_NAME")

	http.HandleFunc("/metrics/", func(w http.ResponseWriter, rq *http.Request) {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprintf("tick_random_metric{application=\"%s\"} %v", app, r.Float32()*100))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprintf("Application: %s", app))
	})

	fmt.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
