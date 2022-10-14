package handlers

import (
	"golang/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{
		l: l,
	}
}
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducs()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable convert to json", http.StatusInternalServerError)
	}
}
