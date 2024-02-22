package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type CartRepo interface {
	AddCart(model.RequestAddCart) (model.ResponseSuccessCart, *resty.Response, error)
	ListCart() ([]model.ResponseCart, *resty.Response, error)
	DetailCart(model.RequesCardById) (model.ResponseCart, *resty.Response, error)
	DeleteCart(model.RequesCardById) (model.ResponseCart, *resty.Response, error)
}

type cartRepo struct {
	client *resty.Client
}

// AddCart implements CartRepo.
func (r cartRepo) AddCart(req model.RequestAddCart) (model.ResponseSuccessCart, *resty.Response, error) {
	var result model.ResponseSuccessCart
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s%s", "http://", os.Getenv("SERVICE_HOST_CART"), "/api/add"))
		//Post("http://localhost:8087/api/add")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err

}

// DeleteCart implements CartRepo.
func (r cartRepo) DeleteCart(req model.RequesCardById) (model.ResponseCart, *resty.Response, error) {
	var result model.ResponseCart
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"cart_id": req.Id,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s%s", "http://", os.Getenv("SERVICE_HOST_CART"), "/api/detail-lps"))
		//Delete("http://localhost:8085/api/delete")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

// DetailCart implements CartRepo.
func (r cartRepo) DetailCart(req model.RequesCardById) (model.ResponseCart, *resty.Response, error) {
	var result model.ResponseCart
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"cart_id": req.Id,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s%s", "http://", os.Getenv("SERVICE_HOST_CART"), "/api/detail-lps"))
		//Get("http://localhost:8085/api/detail")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

// ListCart implements CartRepo.
func (r cartRepo) ListCart() ([]model.ResponseCart, *resty.Response, error) {
	var result []model.ResponseCart
	resp, err := r.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s%s", "http://", os.Getenv("SERVICE_HOST_CART"), "/api/detail-lps"))
		//Get("http://localhost:8085/api/list")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

func NewCartRepo() CartRepo {
	return cartRepo{
		resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
