package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/newrelic/go-agent/v3/newrelic"
	log "github.com/sirupsen/logrus"
)

func fromXForwardedFor(x_forwarded_for string) string {
	return strings.TrimSpace(strings.Split(x_forwarded_for, ",")[0])
}

func fromRemoteAddr(remote_addr string) string {
	return strings.Split(remote_addr, ":")[0]
}

func handler(w http.ResponseWriter, r *http.Request) {
	var remote_addr string

	x_forwarded_for := r.Header.Get("X-Forwarded-For")
	if x_forwarded_for != "" {
		remote_addr = fromXForwardedFor(x_forwarded_for)
	} else {
		remote_addr = fromRemoteAddr(r.RemoteAddr)
	}

	fmt.Fprintf(w, remote_addr)
}

func main() {
	configuration := "0.0.0.0:8000"

	useNewRelic := flag.Bool("use-newrelic", false, "Turn on newrelic")
	flag.Parse()
	if flag.NArg() > 0 {
		configuration = flag.Args()[0]
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEWRELIC_APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEWRELIC_CONFIG_LICENSE")),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(config *newrelic.Config) {
			config.Enabled = *useNewRelic
		},
	)

	if err != nil {
		log.Error("Couldn't initialize New Relic")
		os.Exit(1)
	}

	log.Info("Starting server listening to: ", configuration)
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", handler))
	http.ListenAndServe(configuration, nil)
}
