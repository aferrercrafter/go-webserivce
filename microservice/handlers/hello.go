package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello Handler
type Hello struct {
	l *log.Logger
}

// NewHello Hello constructor
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHttp serve
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
