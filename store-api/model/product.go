package model

import (
	"blog-api/store"
)

type Product struct {
	ID              int
	Name            string
	Path            string
	Description     string
	Tags            []string
	Price           int
	Preview         string
	Characteristics string
}

func NewProduct(a store.Product) Product {
	return Product{
		ID:              a.ID,
		Name:            a.Name,
		Path:            a.Path,
		Description:     a.Description,
		Tags:            a.Tags,
		Preview:         a.Preview.String,
		Price:           a.Price,
		Characteristics: a.Characteristics,
	}
}

func NewProducts(a []store.Product) []Product {
	var articles []Product
	for i := range a {
		articles = append(articles, NewProduct(a[i]))
	}
	return articles
}
