package model

type RequestAddCart struct {
	CustomerID uint `json:"customer_id"`
	ProductID  uint `json:"product_id"`
	Quantity   uint `json:"qty"`
}

type RequesCardById struct {
	Id string `form:"cart_id"`
}

type ResponseSuccessCart struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  uint    `json:"qty"`
	Msq       string  `json:"msg"`
}

type ResponseCart struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"total_amount"`
	Quantity  uint    `json:"qty"`
}
