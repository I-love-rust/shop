package purchase

import (
	"blog-api/config"
	"blog-api/rest/req"
	"blog-api/store"
	"blog-api/tools/tokenmanager"
)

type Purchase interface {
	SuccessPurchase(ctx *req.Ctx) error
	GenerateBill(ctx *req.Ctx) error
}

type Service struct {
	purchaseStore store.PurchaseStore
	productStore  store.ProductStore
	tokenManager  tokenmanager.Tool
	cashbox       config.Cashbox
}

func New(purchaseStore store.PurchaseStore, token tokenmanager.Tool, productStore store.ProductStore, cashbox config.Cashbox) *Service {
	return &Service{
		purchaseStore: purchaseStore,
		productStore:  productStore,
		tokenManager:  token,
		cashbox:       cashbox,
	}
}
