package response

type SuccessAddCustomer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required,email`
}

type LoginResponse struct {
	Username string `json:"username"`
	Msg      string `json:messege`
}
