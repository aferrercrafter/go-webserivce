package handlers

import (
	"log"
	"net/http"
)

// Goodbye Handler
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye Goodbye constructor
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// ServeHttp serve
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeee"))
}
