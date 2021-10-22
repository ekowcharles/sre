package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func logRequest(w *http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}

	logf(Debug, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		logf(Debug, "%q: %q\n", k, v)
	}

	logf(Debug, "Host = %q\n", r.Host)
	logf(Debug, "RemoteAddr = %q\n", r.RemoteAddr)
	logf(Debug, "Accept = %q", r.Header["Accept"])
}

func setHeaders(w *http.ResponseWriter, req *http.Request, url string) {
	(*w).Header().Set("Content-type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", url)
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Cache, X-CSRF-Token, Authorization, token")
}

func buildResponse(w http.ResponseWriter, c int) {
	r := Payload{strconv.Itoa(c), http.StatusText(c)}
	JSON, _ := json.Marshal(r)

	w.WriteHeader(c)
	io.WriteString(w, string(JSON))
}

func wrapper(h http.HandlerFunc, as *AppSetting) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logln(Debug, "============ BEGIN ============")
		logRequest(&w, r)
		defer logln(Debug, "============  END  =============")

		logln(Debug, "Logging general headers")
		setHeaders(&w, r, as.LoadBalancerUrl)

		if r.Method == http.MethodOptions {
			logln(Debug, "Exiting - OPTIONS request")
			return
		}

		logln(Debug, "Trigering original request")
		h.ServeHTTP(w, r)
	})
}
