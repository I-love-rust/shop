package purchase

import (
	"blog-api/pkg/errs"
	"blog-api/rest/req"
	"blog-api/store"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func (s *Service) SuccessPurchase(ctx *req.Ctx) error {
	user := ctx.Request.URL.Query().Get("user")
	product := ctx.Request.URL.Query().Get("product")
	price := ctx.Request.URL.Query().Get("price")
	transHashSum := ctx.Request.URL.Query().Get("ts")

	realHashSum := GetMD5Hash(fmt.Sprintf("%s,%s,%s,%s", user, product, price, s.cashbox.OutHash))
	if realHashSum != transHashSum {
		return errs.ReturnError(ctx.Writer, errs.InvalidHashsum)
	}

	userID, err := strconv.Atoi(user)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InvalidUrlParam)
	}
	productID, err := strconv.Atoi(product)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InvalidUrlParam)
	}

	err = s.purchaseStore.RegisterPurchase(store.RegisterPurchaseOpts{
		UserID:    userID,
		ProductID: productID,
	})

	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InternalServerError)
	}

	return ctx.ReturnSuccessNil()
}

func (s *Service) GenerateBill(ctx *req.Ctx) error {
	productStr := ctx.Request.URL.Query().Get("product")

	productID, err := strconv.Atoi(productStr)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InvalidUrlParam)
	}

	product, err := s.productStore.GetProductByID(productID)

	user, err := s.tokenManager.ValidateJWTToken(ctx.BearerToken())
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.AccessTokenHasExpired)
	}

	URL, err := url.Parse(s.cashbox.ServerAddr)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InternalServerError)
	}

	urlQuery := URL.Query()

	urlQuery.Add("user", fmt.Sprint(user.UserID))
	urlQuery.Add("product", fmt.Sprint(product.ID))
	urlQuery.Add("price", fmt.Sprint(product.Price))
	urlQuery.Add("ts",
		GetMD5Hash(fmt.Sprintf("%d,%d,%d,%s", user.UserID, product.ID, product.Price, s.cashbox.InHash)))

	URL.RawQuery = urlQuery.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL.String(), nil)

	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InternalServerError)
	}
	_, err = client.Do(req)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InternalServerError)
	}

	return ctx.ReturnSuccessNil()
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
