package middleware

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s from %s %v\n", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
}

func NewLogger(handler http.Handler) http.Handler {
	return &Logger{
		handler: handler,
	}
}
