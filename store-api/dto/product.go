package dto

import (
	"blog-api/model"
)

type Product struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Path            string   `json:"path"`
	Description     string   `json:"description"`
	Tags            []string `json:"tags"`
	Price           int      `json:"price"`
	Preview         string   `json:"preview"`
	Characteristics string   `json:"characteristics"`
}

func NewProduct(a model.Product) Product {
	return Product{
		ID:              a.ID,
		Name:            a.Name,
		Path:            a.Path,
		Description:     a.Description,
		Tags:            a.Tags,
		Preview:         a.Preview,
		Price:           a.Price,
		Characteristics: a.Characteristics,
	}
}

func NewProducts(a []model.Product) []Product {
	articles := []Product{}
	for i := range a {
		articles = append(articles, NewProduct(a[i]))
	}
	return articles
}
