package grpchandlers

import (
	"context"
	"simple-commerce/internal/database"
	"simple-commerce/internal/models"
	productspb "simple-commerce/internal/protos"
	"time"

	"gorm.io/gorm"
)

type ProductGRPCServer struct {
	productspb.UnimplementedProductServiceServer // promoted field
	db                                           *gorm.DB
}

func NewProductGRPCServer(db *gorm.DB) *ProductGRPCServer {
	return &ProductGRPCServer{db: db}
}

func (pgs *ProductGRPCServer) CreateProduct(ctx context.Context, req *productspb.CreateProductRequest) (*productspb.CreateProductResponse, error) {

	productDB := database.NewProductDB(pgs.db)

	product := new(models.Product)

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = int(req.Stock)
	product.Status = req.Status
	product.LastModified = time.Now().Unix()

	if err := product.Validate(); err != nil {
		return nil, err
	}

	product, err := productDB.Create(product)
	if err != nil {
		return nil, err
	}

	res := new(productspb.CreateProductResponse)
	res.Id = uint64(product.ID)
	res.Name = product.Name
	res.Description = product.Description
	res.Stock = int32(product.Stock)
	res.Price = product.Price
	res.Status = product.Status
	res.LastModified = product.LastModified
	return res, nil
}
