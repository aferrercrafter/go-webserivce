package handlers

import (
	"log"
	"net/http"

	"github.com/aferrercrafter/microservice-go/data"
)

// Products Handler
type Products struct {
	l *log.Logger
}

// NewProducts Products constructor
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHttp serve
func (h *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
	}
}
