package dto

import (
	"github.com/gofiber/fiber/v2"
)

// Base Response DTO
type MetaPagination struct {
	CurrentPage int64 `json:"current_page"`
	NextPage    int64 `json:"next_page"`
	PrevPage    int64 `json:"prev_page"`
	TotalPages  int64 `json:"total_pages"`
	TotalCount  int64 `json:"total_count"`
}

func CreatePaginationResponse(total int64, page int, limit int) *MetaPagination {
	return &MetaPagination{
		CurrentPage: int64(page),
		NextPage:    int64(page + 1),
		PrevPage:    int64(page - 1),
		TotalPages:  (total + int64(limit) - 1) / int64(limit),
		TotalCount:  total,
	}
}

type BaseResponse struct {
	Success    bool            `json:"success"`
	Message    string          `json:"message,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
	Error      string          `json:"error,omitempty"`
	Pagination *MetaPagination `json:"pagination,omitempty"`
}

// Pagination DTO
type PaginationRequest struct {
	Limit int `query:"limit" validate:"min=1"`
	Page  int `query:"page" validate:"min=1,max=100"`
}

func (p *PaginationRequest) GetLimit() int {
	if p.Limit == 0 {
		return 10
	}
	return p.Limit
}

func (p *PaginationRequest) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p *PaginationRequest) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func JSONWithMessage(c *fiber.Ctx, status int, message string, data interface{}, pagination *MetaPagination) error {
	return c.Status(status).JSON(BaseResponse{
		Success:    status < 400,
		Message:    message,
		Pagination: pagination,
		Data:       data,
	})
}

func ErrorWithMessage(c *fiber.Ctx, status int, message string, err error, pagination *MetaPagination) error {
	return c.Status(status).JSON(BaseResponse{
		Success:    false,
		Message:    message,
		Pagination: pagination,
		Error:      err.Error(),
	})
}
