package http

// CreateUserRequest es lo que el cliente nos envía en el POST
type CreateUserRequest struct {
	Nombre string  `json:"nombre" binding:"required"`
	Email  string  `json:"email" binding:"required,email"`
	Edad   float64 `json:"edad" binding:"gte=0"`
}

// UserResponse es lo que nosotros devolvemos en el GET
type UserResponse struct {
	ID     int64   `json:"id"`
	Nombre string  `json:"nombre"`
	Email  string  `json:"email"`
	Edad   float64 `json:"edad"`
}
