package main

import (
	"fmt"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

var sysLog, sysLogErr = syslog.Dial("udp", "192.168.33.10:514", syslog.LOG_WARNING|syslog.LOG_DAEMON, "pp")

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()

	fmt.Fprintf(sysLog, "Increasing counter ...")
	counter++
	sysLog.Debug("Counter increased")

	fmt.Fprintf(w, fmt.Sprintf("pong %d", counter))
	mutex.Unlock()
}

func main() {
	if sysLogErr != nil {
		log.Fatal(sysLogErr)
	}

	port := os.Getenv("PING_PORT")

	http.HandleFunc("/ping", incrementCounter)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
