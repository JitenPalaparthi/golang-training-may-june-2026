package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"simple-commerce/internal/handlers"
	"simple-commerce/internal/models"
	"simple-commerce/internal/repositories/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProducthandler_CreateProduct_Success(t *testing.T) {

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIProductDB(ctrl)

	input := &models.Product{
		Name:        "Laptop",
		Description: "Some Laptop",
		Price:       50000,
	}

	mockRepo.EXPECT().Create(gomock.Any()).Return(input, nil).Times(1)

	handler := handlers.NewProductHandler(mockRepo)

	body := `{
	"name":"Laptop",
	"description":"Some Laptop",
	"price":50000
	}`

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	handler.CreateProduct(ctx)

	// if http.StatusCreated != rec.Code {
	// 	t.Errorf("expeected 201 but got:%d", rec.Code)
	// }

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "Laptop")

}

func TestProducthandler_CreateProduct_ValidateFailure(t *testing.T) {

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIProductDB(ctrl)

	//mockRepo.EXPECT().Create(gomock.Any()).Return(nil, errors.New("invalid product description")).Times(1)

	//mockRepo.EXPECT().Create(gomock.Any()).Times(1)

	handler := handlers.NewProductHandler(mockRepo)

	body := `{
	"name":"Laptop",
	"price":50000
	}`

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	handler.CreateProduct(ctx)

	// if http.StatusCreated != rec.Code {
	// 	t.Errorf("expeected 201 but got:%d", rec.Code)
	// }

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid product description")

}

// func TestA(t *testing.T) {

// 	num := 10

// 	result := num * num

// 	assert.Equal(t, result, 101)

// 	assert.Equal(t, result+10, 110)

// }

// func TestB(t *testing.T) {
// 	num := 10
// 	result := num * num
// 	//assert.Equal(t, result, 101)
// 	assert.Equal(t, result+10, 110)
// }
