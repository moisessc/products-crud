package service

import (
	"context"
	"fmt"
	"testing"

	"products-crud/internal/model"
	"products-crud/internal/repository/mocks"
	"products-crud/pkg/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductService(t *testing.T) {
	testCases := map[string]struct {
		mockReturn error
		params     struct {
			ctx     context.Context
			product *model.Product
		}
		errorExpected error
	}{
		"could_not_create_product": {
			mockReturn: errors.ErrFailedToSaveProduct,
			params: struct {
				ctx     context.Context
				product *model.Product
			}{
				ctx:     context.Background(),
				product: model.NewProductWithoutId(1, 1, 50, "Macbook 2021", 3200.56, false),
			},
			errorExpected: fmt.Errorf("persistence failed: %w", errors.ErrFailedToSaveProduct),
		},
		"create_product_success": {
			mockReturn: nil,
			params: struct {
				ctx     context.Context
				product *model.Product
			}{
				ctx:     context.Background(),
				product: model.NewProductWithoutId(1, 1, 50, "Macbook 2021", 3200.56, false),
			},
			errorExpected: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockRepository := mocks.ProductRepository{}
			mockRepository.On(
				"Save",
				mock.AnythingOfType("*context.emptyCtx"),
				mock.AnythingOfType("*model.ProductEntity"),
			).Return(tc.mockReturn)

			service := NewProductService(&mockRepository)
			err := service.CreateProduct(tc.params.ctx, tc.params.product)
			if tc.errorExpected != nil {
				assert.EqualError(t, err, tc.errorExpected.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
