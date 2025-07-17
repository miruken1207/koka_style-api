package models

type SignUp_input struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login_input struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
