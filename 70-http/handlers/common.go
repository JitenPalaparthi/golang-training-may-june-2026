package handlers

import (
	"log/slog"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		slog.Warn("called on not implemented http method" + r.Method)
		return
	}
	slog.Debug("pong")
	w.Write([]byte("pong"))
}

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		slog.Warn("called on not implemented http method" + r.Method)
		return
	}
	slog.Debug("ok")
	w.Write([]byte("ok"))
}
