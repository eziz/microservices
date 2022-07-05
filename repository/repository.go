package repository

import (
	"accelebrate.com/svc/models"
)

var products []models.Product
var nextID = 1

func GetProducts() []models.Product{
	return products
}

func AddProduct(product models.Product) int{
	product.ID=nextID
	nextID++
	product=append(products,product)
	return product.ID
}