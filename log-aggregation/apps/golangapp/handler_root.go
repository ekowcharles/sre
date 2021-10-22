package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) root(w http.ResponseWriter, rq *http.Request) {
	JSON, _ := json.Marshal(h.as)

	w.WriteHeader(http.StatusOK)

	io.WriteString(w, string(JSON))
}
