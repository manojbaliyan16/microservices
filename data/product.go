package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdateOn    string
	DeletionOn  string
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milky Coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong Coffee without milk",
		Price:       1.99,
		SKU:         "def1234",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
