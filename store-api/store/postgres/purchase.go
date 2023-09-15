package postgres

import (
	"blog-api/store"
	"database/sql"
	"log"
)

type PurchaseStore struct {
	db *sql.DB
}

func NewPurchaseStore(db *sql.DB) store.PurchaseStore {
	return &PurchaseStore{
		db: db,
	}
}

func (s *PurchaseStore) RegisterPurchase(opts store.RegisterPurchaseOpts) error {
	_, err := s.db.Exec("CALL register_purchase($1, $2)", opts.UserID, opts.ProductID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
