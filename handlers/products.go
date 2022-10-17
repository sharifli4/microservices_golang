// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta

package handlers

import (
	"fmt"
	"golang.org/x/net/context"
	"golang/data"
	"log"
	"net/http"
)
// A list of products returns in the response
// swagger:response productsResponseWrapper
type productsResponseWrapper struct {
	// All products in the system
	// in:body
	Body []data.Product

}
// swagger:response noContent
type productsNoContent struct {}
//swagger:parameters DeleteProduct
type productIDParameterWrapper struct{
	// the id of the product to delete from the database
	// in:path
	// required:true
	ID int `json:"id"`
}
// Products  is a http.Handler
type Products struct {
	l *log.Logger
}
// NewProducts creates a product handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{
		l: l,
	}
}
type KeyProduct struct{}
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product",err)
			http.Error(rw, "Unable to unmarshall json", http.StatusInternalServerError)
		}
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product",err)
			http.Error(
				rw,
				fmt.Sprintf("Unable to validate product: %s",err),
				http.StatusBadRequest,
				)
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
