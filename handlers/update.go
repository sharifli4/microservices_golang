package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang/data"
	"net/http"
	"strconv"
)
// UpdateProduct update a product handler with the given logger
func (p Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(mux.Vars(r))
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Cannot convert id to int", http.StatusInternalServerError)
	}
	p.l.Println("Handle PUT Request", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFind {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
}