package handlers

import (
	"encoding/json"
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
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(d)
}
