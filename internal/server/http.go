package server

import (
	log "log/slog"
	"net/http"
)

const port = "8080" // TODO: Replace with env variable
const endpoint = ":" + port

func New() *http.Server {
	svr := &http.Server{
		Addr:    endpoint,
		Handler: http.HandlerFunc(handler),
	}
	log.Info("Server listening on port 8080...")
	log.Error("server.New(): %s", svr.ListenAndServeTLS("server.crt", "server.key"))

	return svr
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info("Connected to %s", r.Proto)
	_, _ = w.Write([]byte("Hello, World!"))
}
