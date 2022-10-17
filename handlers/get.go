package handlers

import (
	"golang/data"
	"net/http"
)
// swagger:route GET /products products listProducts
// Return list of products
// responses:
//	200: productsResponse
// GetProducts returns the products  from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handled request")
	lp := data.GetProducts()
	w.Header().Set("content-Type", "application/json")
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable convert to json", http.StatusInternalServerError)
	}
}