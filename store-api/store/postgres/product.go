package postgres

import (
	"blog-api/store"
	"database/sql"

	"github.com/lib/pq"
)

type ProductStore struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) store.ProductStore {
	return &ProductStore{
		db: db,
	}
}

func (s *ProductStore) CreateProduct(opts store.CreateProductOpts) error {
	_, err := s.db.Exec("CALL create_product($1, $2, $3, $4, $5, $6, $7)",
		opts.Name, opts.Path, opts.Description, pq.Array(opts.Tags), opts.Price, opts.Preview, opts.Characteristics)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductStore) GetProductByPath(path string) (store.Product, error) {
	var product store.Product

	res := s.db.QueryRow("SELECT * FROM get_product_by_path($1)", path)

	err := res.Scan(&product.ID, &product.Name, &product.Path, &product.Description, pq.Array(&product.Tags), &product.Price, &product.Preview, &product.Characteristics)
	if err != nil {
		return store.Product{}, err
	}

	return product, nil
}

func (s *ProductStore) GetProductByID(id int) (store.Product, error) {
	var product store.Product

	res := s.db.QueryRow("SELECT * FROM get_product_by_id($1)", id)

	err := res.Scan(&product.ID, &product.Name, &product.Path, &product.Description, pq.Array(&product.Tags), &product.Price, &product.Preview, &product.Characteristics)
	if err != nil {
		return store.Product{}, err
	}

	return product, nil
}

func (s *ProductStore) GetProductForFeed(page int) ([]store.Product, error) {
	var products []store.Product

	res, err := s.db.Query("SELECT * FROM get_products_page($1)", page)
	if err != nil {
		return []store.Product{}, err
	}

	for res.Next() {
		var product store.Product
		err := res.Scan(&product.ID, &product.Name, &product.Path, &product.Description, pq.Array(&product.Tags), &product.Price, &product.Preview, &product.Characteristics)
		if err != nil {
			return []store.Product{}, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (s *ProductStore) SearchProduct(opts store.SearchProductOpts) ([]store.Product, error) {
	var products []store.Product

	res, err := s.db.Query("SELECT * FROM search_product($1, $2, $3)", opts.Page, pq.Array(opts.Tags), opts.Text)
	if err != nil {
		return []store.Product{}, err
	}

	for res.Next() {
		var product store.Product
		err := res.Scan(&product.ID, &product.Name, &product.Path, &product.Description, pq.Array(&product.Tags), &product.Price, &product.Preview, &product.Characteristics)
		if err != nil {
			return []store.Product{}, err
		}

		products = append(products, product)
	}

	return products, nil
}
