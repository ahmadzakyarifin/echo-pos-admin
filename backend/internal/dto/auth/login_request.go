package auth

type LoginRequest struct {
	Username string `json:"username" binding:"required_without=Email"`
	Email    string `json:"email" binding:"required_without=Username"`
	Password string `json:"password" binding:"required min=6"`
}