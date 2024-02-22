package request

// update user isactive
type AddCustomer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required,email`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
