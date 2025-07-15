package models

// @Schema(description="User registration data")
// example={"email": "john@example.com", "username": "john_doe", "password": "secret"}
type SignUp_input struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Schema(description="User login credentials")
// example={"username": "john_doe", "password": "secret"}
type Login_input struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
	