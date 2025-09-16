package repository

import "gorm.io/gorm"

type Repositories struct {
	User    UserRepository
	Product ProductRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:    NewUserRepository(db),
		Product: NewProductRepository(db),
		// Order:   NewOrderRepository(db),
	}
}
