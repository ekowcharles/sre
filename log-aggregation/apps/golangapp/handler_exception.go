package main

import (
	"net/http"
)

func (h *Handler) exception(w http.ResponseWriter, rq *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	panic("Doing this only for the logs")
}
