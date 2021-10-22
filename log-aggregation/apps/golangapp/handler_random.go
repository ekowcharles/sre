package main

import (
	"math/rand"
	"net/http"
)

func (h *Handler) random(w http.ResponseWriter, rq *http.Request) {
	i := rand.Intn(sl)

	buildResponse(w, ss[i])
}
