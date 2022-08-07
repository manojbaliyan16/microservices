package handlers

import (
	"encoding/json"
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
	lp := data.GetProducts()       // here Lp is having the list of product now how can we return it to the user The way is to convert it(LP) into JSON can be done by using encoding JSON
	bytes, err := json.Marshal(lp) // bytes--> slice of data
	if err != nil {
		http.Error(w, "unable to marshal Json", http.StatusInternalServerError)
	}
	w.Write(bytes)
	/*d, err := ioutil.ReadAll(r.Body)
	if err!=nil{

	}*/
}
