package response

import (
	"time"
)

// User Response DTOs
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// Address   UserAddressOutput `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// type UserAddressOutput struct {
// 	Street string `json:"street"`
// 	City   string `json:"city"`
// }

// type UsersResponse struct {
// 	Users []UserResponse `json:"users"`
// 	Total int64          `json:"total"`
// 	Page  int            `json:"page"`
// 	Size  int            `json:"size"`
// }
