package models

import "errors"

type Product struct {
	// gorm.Model
	CommonModel
	Name        string  `json:"name" gorm:"size:200;not null"`
	Description string  `json:"description" gorm:"not null;default:some product"`
	Price       float32 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock"`
	Status      string  `json:"status" gorm:"default:active"`
}

type CommonModel struct {
	ID uint `gorm:"primarykey"`
	// We explicitly cast the epoch to a bigint so Postgres drops the decimal point
	LastModified int64 `json:"last_modified" gorm:"column:last_modified;type:bigint;not null;default:extract(epoch from now())::bigint"`
}

func (p *Product) Validate() error {

	if p.Name == "" {
		return errors.New("invalid product name")
	}

	if p.Description == "" {
		return errors.New("invalid product description ")
	}

	if p.Price <= 0 {
		return errors.New("invalid product price.Price must be greater than 0")
	}

	return nil

}
