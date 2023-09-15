package product

import (
	"blog-api/pkg/errs"
	"blog-api/rest/req"
	"blog-api/store"
	"blog-api/tools/chars"
	"blog-api/tools/tokenmanager"
)

type NewArticleRequest struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Tags            []string `json:"tags"`
	Price           int      `json:"price"`
	Preview         string   `json:"preview"`
	Characteristics string   `json:"characteristics"`
}

type NewArticleResponse struct {
	Path string `json:"path"`
}

func (s *Service) NewProduct(ctx *req.Ctx) error {
	var r NewArticleRequest

	err := ctx.ParseJSON(&r)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InvalidJSON)
	}

	pathSalt, err := tokenmanager.GenerateRandomSalt(2)
	if err != nil {
		return err
	}

	path := chars.ToLatin(r.Name) + "-" + pathSalt

	err = s.productStore.CreateProduct(store.CreateProductOpts{
		Name:            r.Name,
		Path:            path,
		Description:     r.Description,
		Tags:            r.Tags,
		Preview:         r.Preview,
		Price:           r.Price,
		Characteristics: r.Characteristics,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(NewArticleResponse{Path: path})
}
