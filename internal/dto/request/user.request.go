package request

import "clean_architecture/internal/dto"

// User Request DTOs
type ListUserRequest struct {
	Search   string                `json:"search" validate:"omitempty,max=100"`
	Paginate dto.PaginationRequest `json:"paginate"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"omitempty,min=3,max=50"`
	Email    string `json:"email" validate:"omitempty,email"`
}
