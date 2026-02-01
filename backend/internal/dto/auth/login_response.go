package auth

type LoginSuccessResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
	Role   string `json:"role"`
}

type LoginErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}