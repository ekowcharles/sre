package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	defer logln(Debug, "Exited!")

	rand.Seed(time.Now().UnixNano())

	port := getEnv("PORT", "8040")
	as := &AppSetting{getEnv("APP_VERSION", "0.0.0"), getEnv("LOAD_BALANCER_URL", "")}
	h := Handler{as: as}

	http.HandleFunc("/", wrapper(h.root, as))
	http.HandleFunc("/http/", wrapper(h.http, as))
	http.HandleFunc("/random", wrapper(h.random, as))
	http.HandleFunc("/exception", wrapper(h.exception, as))

	logf(Debug, "Listening on port :%s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
