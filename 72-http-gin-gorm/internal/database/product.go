package database

import (
	"log/slog"
	"simple-commerce/internal/models"
	"simple-commerce/internal/repositories"

	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) repositories.IProductDB {
	return &ProductDB{db}
}

func (p *ProductDB) Create(product *models.Product) (*models.Product, error) {
	if p.DB == nil {
		panic("nil database")
	}

	// err := p.DB.AutoMigrate(&models.Product{}) // if no table this table would be created
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	return nil, err
	// }
	tx := p.DB.Create(product)
	if tx.Error != nil {
		slog.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return product, nil
}
