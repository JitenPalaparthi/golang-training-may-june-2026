package handlers

import (
	"log/slog"
	"net/http"
	"simple-commerce/internal/models"
	"simple-commerce/internal/repositories"

	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	CreateProduct(ctx *gin.Context)
}

type ProductHandler struct {
	repositories.IProductDB //promoted so all methods automatically apprear
}

func NewProductHandler(iproductdb repositories.IProductDB) IProductHandler {
	return &ProductHandler{iproductdb}
}

func (p *ProductHandler) CreateProduct(ctx *gin.Context) {
	var product *models.Product

	if p.IProductDB == nil {
		slog.Error("nil IProductDB object")
		ctx.String(http.StatusBadRequest, "oops! something went wrong, try again or consult admin")
		ctx.Abort()
		return
	}

	err := ctx.Bind(&product) // unmarshal
	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, "oops! something went wrong, try again or consult admin")
		ctx.Abort()
		return
	}

	err = product.Validate()

	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	//product, err = p.IProductDB.Create(product) // This is called on an interface
	product, err = p.Create(product) // This is called on an interface
	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, "oops! something went wrong, try again or consult admin")
		ctx.Abort()
		return
	}

	ctx.JSON(201, product)

}

// func prepare(table string, pair map[string]string) string {

// 	query := fmt.Sprint("Select * from where ", table)

// 	for k, v := range pair {
// 		query += fmt.Sprint(k, "'", v, "'")
// 	}
// 	return query

// }
