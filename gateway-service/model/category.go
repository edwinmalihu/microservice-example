package model

type RequestAdd struct {
	Category string `json:"category"`
}

type RequesListCategoryById struct {
	Id string `form:"category_id"`
}

type RequestUpdateCategory struct {
	Id       int    `json:"category_id"`
	Category string `json:"category"`
}

type ResponseSuccess struct {
	Category string `json:"category"`
	Msg      string `json:"msg"`
}

type ResponseCategory struct {
	Id       uint   `json:"category_id"`
	Category string `json:"category"`
}
