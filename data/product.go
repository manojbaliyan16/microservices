package data

import "time"

type Product struct {
	ID          int     `json:"id"` // JSON as struct tag, add the JSON annotation, It will help us to change the better formatting like change the field name, omit the field or completely omit if it is completely empty
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"_"`
	UpdateOn    string  `json:"_"`
	DeletionOn  string  `json:"_"`
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
