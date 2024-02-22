package request

type RequestAddProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stok        uint    `json:"stok"`
	CategoryID  uint    `json:"category_id"`
}
type RequestUpdateProduct struct {
	Id          uint    `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stok        uint    `json:"stok"`
	CategoryID  uint    `json:"category_id"`
}

type RequesByIdProduct struct {
	Id string `form:"product_id"`
}
type RequesByIdCategory struct {
	Id string `form:"category_id"`
}
