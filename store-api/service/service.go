package service

import (
	"blog-api/config"
	"blog-api/service/product"
	"blog-api/service/purchase"
	"blog-api/service/user"
	"blog-api/store"
	"blog-api/tools/tokenmanager"
)

type Service struct {
	User     user.User
	Product  product.Product
	Purchase purchase.Purchase
}

type Opts struct {
	UserStore     store.UserStore
	SessionStore  store.SessionStore
	ProductStore  store.ProductStore
	PurchaseStore store.PurchaseStore
	Token         tokenmanager.Tool
	Cashbox       config.Cashbox
	AssetsDir     string
}

func New(s Opts) *Service {
	return &Service{
		User:     user.New(s.UserStore, s.SessionStore, s.Token),
		Product:  product.New(s.UserStore, s.ProductStore, s.Token, s.AssetsDir),
		Purchase: purchase.New(s.PurchaseStore, s.Token, s.ProductStore, s.Cashbox),
	}
}
