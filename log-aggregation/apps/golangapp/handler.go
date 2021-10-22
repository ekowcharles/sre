package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func logRequest(w *http.ResponseWriter, r *http.Request) {
	var sb strings.Builder
	for k, v := range r.Header {
		if k == "Authorization" || k == "Proxy-Authorization" {
			sb.WriteString(fmt.Sprintf(" %q: [*****]", k))
		} else {
			sb.WriteString(fmt.Sprintf(" %q: %q", k, v))
		}
	}

	logf(Debug, "%s %s %s %s %s%s", r.Method, r.URL, r.Proto, r.Host, r.RemoteAddr, sb.String())
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
