package repositories

import "simple-commerce/internal/models"

//go:generate mockgen -source=repository.go -destination=mocks/db_mock.go -package=mocks

type IProductDB interface {
	Create(product *models.Product) (*models.Product, error)
}

type IOrderDB interface {
	Create(order *models.Order) (*models.Order, error)
}
