package handler

import "clean_architecture/internal/service"

type Handlers struct {
	User *UserHandler
}

// func NewHandlers(service)
func NewHandlers(service *service.Services) *Handlers {
	return &Handlers{
		User: NewUserHandler(service),
	}
}
