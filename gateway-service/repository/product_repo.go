package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type ProductRepo interface {
	AddProduct(model.RequestAddProduct) (model.ResponseSucessProduct, *resty.Response, error)
	UpdateProduct(model.RequestUpdateProduct) (model.ResponseSucessProduct, *resty.Response, error)
	DetailProduct(model.RequesByIdProduct) (model.ResponseDetailProduct, *resty.Response, error)
	ListProduct() ([]model.ResponseDetailProduct, *resty.Response, error)
	ListProductByCategory(model.RequesByIdCategory) ([]model.ResponseDetailProduct, *resty.Response, error)
}

type productRepo struct {
	client *resty.Client
}

// AddProduct implements ProductRepo.
func (r productRepo) AddProduct(req model.RequestAddProduct) (model.ResponseSucessProduct, *resty.Response, error) {
	var result model.ResponseSucessProduct
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s%s", "http://", os.Getenv("SERVICE_HOST_PRODUCT"), "/api/add"))
		//Post("http://localhost:8086/api/add")

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

// DetailProduct implements ProductRepo.
func (r productRepo) DetailProduct(req model.RequesByIdProduct) (model.ResponseDetailProduct, *resty.Response, error) {
	var result model.ResponseDetailProduct
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"product_id": req.Id,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s%s", "http://", os.Getenv("SERVICE_HOST_PRODUCT"), "/api/detail-lps"))
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

// ListProduct implements ProductRepo.
func (r productRepo) ListProduct() ([]model.ResponseDetailProduct, *resty.Response, error) {
	var result []model.ResponseDetailProduct
	resp, err := r.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s%s", "http://", os.Getenv("SERVICE_HOST_PRODUCT"), "/api/detail-lps"))
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

// ListProductByCategory implements ProductRepo.
func (r productRepo) ListProductByCategory(req model.RequesByIdCategory) ([]model.ResponseDetailProduct, *resty.Response, error) {
	var result []model.ResponseDetailProduct
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"category_id": req.Id,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s%s", "http://", os.Getenv("SERVICE_HOST_PRODUCT"), "/api/detail-lps"))
		//Get("http://localhost:8085/api/list-ProductByCategory")

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

// UpdateProduct implements ProductRepo.
func (r productRepo) UpdateProduct(req model.RequestUpdateProduct) (model.ResponseSucessProduct, *resty.Response, error) {
	var result model.ResponseSucessProduct
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s%s", "http://", os.Getenv("SERVICE_HOST_PRODUCT"), "/api/add"))
		//Post("http://localhost:8086/api/update")

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

func NewProductRepo() ProductRepo {
	return productRepo{
		resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
