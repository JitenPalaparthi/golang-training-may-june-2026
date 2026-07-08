package repositories

import "simple-commerce/internal/models"

type IProductDB interface {
	Create(product *models.Product) (*models.Product, error)
}
