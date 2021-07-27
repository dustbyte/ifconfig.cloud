package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func fromXForwardedFor(x_forwarded_for string) string {
	return strings.TrimSpace(strings.Split(x_forwarded_for, ",")[0])
}

func handler(w http.ResponseWriter, r *http.Request) {
	var remote_addr string
	uu_id := uuid.NewString()
	logger := log.WithFields(log.Fields{"uuid": uu_id})

	logger.Info("started")
	x_forwarded_for := r.Header.Get("X-Forwarded-For")
	if x_forwarded_for != "" {
		logger.Info("X-Forwarded-For: ", x_forwarded_for)
		remote_addr = fromXForwardedFor(x_forwarded_for)
	} else {
		logger.Info(" RemoteAddr: ", remote_addr)
		remote_addr = strings.Split(r.RemoteAddr, ":")[0]
	}

	logger.Info("replying to ", remote_addr)
	fmt.Fprintf(w, remote_addr)
	logger.Info("terminated")
}

func main() {
	configuration := "0.0.0.0:8000"
	log.SetFormatter(&log.JSONFormatter{})

	flag.Parse()
	if flag.NArg() > 0 {
		configuration = flag.Args()[0]
	}

	log.Info("Starting server listening to: ", configuration)
	http.HandleFunc("/", handler)
	http.ListenAndServe(configuration, nil)
}
