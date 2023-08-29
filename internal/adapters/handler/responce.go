package handler

type badResponse struct {
	Message string `json:"message"`
}

type successResponse struct {
	Success bool `json:"success"`
}