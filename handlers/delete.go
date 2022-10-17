package handlers

import (
	"github.com/gorilla/mux"
	"golang/data"
	"net/http"
	"strconv"
)

// swagger:route DELETE /products/{id} products DeleteProduct
// Return list of products
// responses:
//	201: noContent
// DeleteProduct deletes the product  from the data store
func (p *Products) DeleteProduct(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("Handle product id ",id)
	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFind {
		http.Error(w,"Product not find",http.StatusNotFound)
	}
	if err != nil {
		http.Error(w,"Product not find",http.StatusNotFound)
	}
}