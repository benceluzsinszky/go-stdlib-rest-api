package middlewares

import (
	"log"
	"net/http"
)

type Logger struct {
	mux *http.ServeMux
}

func NewLogger(mux *http.ServeMux) *Logger {
	return &Logger{mux: mux}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.mux.ServeHTTP(w, r)
	method := r.Method
	path := r.URL.Path
	log.Println(method, path)
}
