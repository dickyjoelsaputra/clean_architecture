package repository

import (
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
