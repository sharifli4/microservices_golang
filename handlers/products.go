package handlers

import (
	"golang/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		p.l.Println(g, r.URL.Path)
		p.l.Println(len(g[0]))
		if len(g) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Cannot convert to int", http.StatusInternalServerError)
		}
		p.updateProduct(id, w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handled request")
	lp := data.GetProducts()
	w.Header().Set("content-Type", "application/json")
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable convert to json", http.StatusInternalServerError)
	}
}
func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST request")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable convert from json", http.StatusInternalServerError)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
func (p Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Request")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshall json", http.StatusInternalServerError)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFind {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
}
