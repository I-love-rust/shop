package store

import (
	"database/sql"
)

type Product struct {
	ID              int
	Name            string
	Path            string
	Description     string
	Tags            []string
	Price           int
	Preview         sql.NullString
	Characteristics string
}

type CreateProductOpts struct {
	Name            string
	Path            string
	Description     string
	Tags            []string
	Price           int
	Preview         string
	Characteristics string
}

type GetProductByPathOpts struct {
	Path string
}

type GetProductFeed struct {
	Page int
}

type SearchProductOpts struct {
	Page int
	Tags interface{}
	Text interface{}
}

type ProductStore interface {
	CreateProduct(opts CreateProductOpts) error
	GetProductByPath(path string) (Product, error)
	GetProductByID(id int) (Product, error)
	SearchProduct(opts SearchProductOpts) ([]Product, error)
	GetProductForFeed(page int) ([]Product, error)
}
