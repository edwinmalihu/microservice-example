package response

type ResponseSuccess struct {
	Category string `json:"category"`
	Msg      string `json:"msg"`
}

type ResponseCategory struct {
	Id       uint   `json:"category_id"`
	Category string `json:"category"`
}
