package model

type Response struct {
	Success     bool        `json:"success"`
	Status      int         `json:"status"`
	ErrorCode   string      `json:"error_code"`
	RespMessage string      `json:"message"`
	RespData    interface{} `json:"data"`
}
