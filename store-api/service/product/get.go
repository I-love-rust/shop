package product

import (
	"blog-api/dto"
	"blog-api/model"
	"blog-api/rest/req"
	"blog-api/store"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (s *Service) GetProduct(ctx *req.Ctx) error {
	path := chi.URLParam(ctx.Request, "path")

	article, err := s.productStore.GetProductByPath(path)
	if err != nil {
		return err
	}

	return ctx.JSON(dto.NewProduct(model.NewProduct(article)))
}

func (s *Service) GetProductForFeed(ctx *req.Ctx) error {
	pageStr := ctx.Request.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return err
	}

	articles, err := s.productStore.GetProductForFeed(page)
	if err != nil {
		return err
	}

	return ctx.JSON(dto.NewProducts(model.NewProducts(articles)))
}

func (s *Service) SearchProduct(ctx *req.Ctx) error {
	tagsStr := ctx.Request.URL.Query().Get("tags[]")
	text := ctx.Request.URL.Query().Get("text")
	pageStr := ctx.Request.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return err
	}

	var tags interface{}
	if toInterface(tagsStr) == nil {
		tags = nil
	} else {
		tags = strings.Split(tagsStr, ",")
	}

	articles, err := s.productStore.SearchProduct(store.SearchProductOpts{
		Page: page,
		Tags: tags,
		Text: toInterface(text),
	})
	if err != nil {
		return err
	}

	return ctx.JSON(dto.NewProducts(model.NewProducts(articles)))
}

func toInterface(param interface{}) interface{} {
	if param == "" {
		return nil
	}
	return param
}
