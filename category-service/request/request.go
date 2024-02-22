package request

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
