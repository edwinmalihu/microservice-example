package repository

import (
	"auth-services/model"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

type OrderRepo interface {
	AddOrder(model.AddOrder) (model.ResponsOrder, *resty.Response, error)
}

type orderRepo struct {
	client *resty.Client
}

// AddOrder implements OrderRepo.
func (r orderRepo) AddOrder(req model.AddOrder) (model.ResponsOrder, *resty.Response, error) {
	var result model.ResponsOrder
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s%s", "http://", os.Getenv("SERVICE_HOST_ORDER"), "/api/add"))
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

func NewOrderRepo() OrderRepo {
	return orderRepo{
		client: resty.New().SetTimeout(3 * 30).SetRetryCount(3),
	}
}
