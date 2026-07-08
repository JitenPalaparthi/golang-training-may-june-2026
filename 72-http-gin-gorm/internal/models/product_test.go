package models_test

import (
	"simple-commerce/internal/models"
	"testing"
)

func TestValidateProduct(t *testing.T) {

	product := new(models.Product) // all are zero values

	err := product.Validate()

	//fmt.Println(err)

	if err == nil { // if there is no error that means the functonality is not properly working so it should be an error
		t.Errorf("expected error , got nil")
	}

	product.Description = "Some product"
	product.Price = 123.123
	product.Stock = 123

	err = product.Validate()

	//fmt.Println(err)

	if err == nil { // if there is no error that means the functonality is not properly working so it should be an error
		t.Errorf("expected error , got nil")
	}

	if err.Error() != "invalid product name" {
		t.Errorf("expected error , got something else")
	}

}

// func TestAbcd(t *testing.T) {
// 	t.Fail()
// }

func TestProductAll(t *testing.T) {

	tests := []struct {
		name    string
		product models.Product
		wantErr bool
	}{

		{
			name: "Valid product name",
			product: models.Product{
				Name:        "Laptop",
				Description: "Good Laptop",
				Price:       12323.23,
			},
			wantErr: false,
		},

		{
			name: "empty product name",
			product: models.Product{
				Name:        "",
				Description: "Good Laptop",
				Price:       12323.23,
			},
			wantErr: true,
		},
		{
			name: "empty description",
			product: models.Product{
				Name:        "Laptop",
				Description: "",
				Price:       12323.23,
			},
			wantErr: true,
		},

		{
			name: "zero price",
			product: models.Product{
				Name:        "Laptop",
				Description: "Decent Laptop",
				Price:       0,
			},
			wantErr: true,
		},
		{
			name: "negative price",
			product: models.Product{
				Name:        "Laptop",
				Description: "Decent Laptop",
				Price:       -1232.123,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.Validate()

			if tt.wantErr && err == nil {
				t.Errorf("expected error , got nil")
			}
			if tt.wantErr == false && err != nil {
				t.Errorf("expected nil , got err %v", err.Error())
			}
		})
	}

}
