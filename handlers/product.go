package handlers

import (
	"log"
	"net/http"

	"microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//need to create a GET request which will return a product define in to product/product.go
	log.Println("Going to handle products of all types ")
	// based on the HTTP request Method take the decision as below using go standard library
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
	}
	// for anyrthinhg else
	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts() // here Lp is having the list of product now how can we return it to the user The way is to convert it(LP) into JSON can be done by using encoding JSON
	// this Lp is the get product
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to Encode", http.StatusInternalServerError)
	}

}
