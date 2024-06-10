package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

const port = "8080" // TODO: Replace with env variable
const endpoint = ":" + port

func New() *http.Server {
	svr := &http.Server{
		Addr:    endpoint,
		Handler: http.HandlerFunc(defaultHandler),
	}
	slog.Info(fmt.Sprintf("Server listening on port %s...", port))
	// TODO: Certs can probably go into an env variable
	// TODO: Use HTTPS in prod?
	// log.Error("server.New(): %s", svr.ListenAndServeTLS("server.crt", "server.key"))
	slog.Error("server.New():", svr.ListenAndServe()) // TODO: If svr errors out, this won't log to console I think

	return svr
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("Connected to %s using %s", endpoint, r.Proto))
	_, _ = w.Write([]byte("Hello, World!"))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
