package main

import (
	"net/http"
	"regexp"
	"strconv"
)

func (h *Handler) http(w http.ResponseWriter, rq *http.Request) {
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
}
