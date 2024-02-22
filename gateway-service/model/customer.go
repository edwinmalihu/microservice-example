package model

// update user isactive
type AddCustomer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required,email"`
}

type SuccessAddCustomer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required,email"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Msg      string `json:"messege"`
	Token    string `json:"token"`
}
