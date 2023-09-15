package product

import (
	"blog-api/rest/req"
	"blog-api/store"
	"blog-api/tools/tokenmanager"
)

type Product interface {
	UploadImage(ctx *req.Ctx) error
	SendFile(ctx *req.Ctx) error
	GetProduct(ctx *req.Ctx) error
	NewProduct(ctx *req.Ctx) error
	GetProductForFeed(ctx *req.Ctx) error
	SearchProduct(ctx *req.Ctx) error
}

type Service struct {
	userStore    store.UserStore
	productStore store.ProductStore
	tokenManager tokenmanager.Tool
	assetsDir    string
}

func New(userStore store.UserStore, articleStore store.ProductStore, token tokenmanager.Tool, assetsDir string) *Service {
	return &Service{
		userStore,
		articleStore,
		token,
		assetsDir,
	}
}
