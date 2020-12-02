package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, fmt.Sprintf("pong %d", counter))
	mutex.Unlock()
}

func main() {
	port := os.Getenv("PING_PORT")

	http.HandleFunc("/ping", incrementCounter)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
