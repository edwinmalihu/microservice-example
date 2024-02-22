package response

type ResponseSucessProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stok        uint    `json:"string"`
	Category    string  `json:"category"`
	Msg         string  `json:"messege"`
}

type ResponseDetailProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stok        uint    `json:"string"`
	Category    string  `json:"category"`
}
