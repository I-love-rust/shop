package store

import (
	"time"
)

type Purchase struct {
	ID      int
	User    int
	Product int
	Time    time.Time
}

type RegisterPurchaseOpts struct {
	UserID    int
	ProductID int
}

type PurchaseStore interface {
	RegisterPurchase(opts RegisterPurchaseOpts) error
}
