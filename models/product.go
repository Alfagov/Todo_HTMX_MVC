package models

import "github.com/google/uuid"

type Product struct {
	Name        string
	UUID        string
	Category    string
	Price       int
	Quantity    int
	Description string
	Status      int
	Image       string
}

func NewProduct(name string, category string, price int, quantity int, description string, status int, image string) Product {
	return Product{name, uuid.New().String(), category, price, quantity, description, status, image}
}

var DefaultProductList = []Product{
	{"Product 1", uuid.New().String(), "Category 1", 100, 10, "Product 1 Description", 2, "https://picsum.photos/200/300"},
	{"Product 2", uuid.New().String(), "Category 2", 200, 20, "Product 2 Description", 0, "https://picsum.photos/200/300"},
	{"Product 3", uuid.New().String(), "Category 3", 300, 30, "Product 3 Description", 2, "https://picsum.photos/200/300"},
}
