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

func TestGetProductsService(t *testing.T) {
	testCases := map[string]struct {
		mockReturn struct {
			err      error
			products []*model.ProductEntity
		}
		ctx              context.Context
		responseExpected []*model.ProductResponse
		errorExpected    error
	}{
		"could_not_get_products": {
			mockReturn: struct {
				err      error
				products []*model.ProductEntity
			}{
				err:      errors.ErrFailedToRetrieveProducts,
				products: nil,
			},
			ctx:           context.Background(),
			errorExpected: fmt.Errorf("failed getting: %w", errors.ErrFailedToRetrieveProducts),
		},
		"get_products_success": {
			mockReturn: struct {
				err      error
				products []*model.ProductEntity
			}{
				err: nil,
				products: []*model.ProductEntity{
					model.NewProductEntity(1, 1, 1, 23, "Macbook Air 2021", 1300.21, false),
				},
			},
			ctx: context.Background(),
			responseExpected: []*model.ProductResponse{
				{
					Id:           1,
					SupplierId:   1,
					CategoryId:   1,
					Stock:        23,
					Name:         "Macbook Air 2021",
					Price:        1300.21,
					Discontinued: false,
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockRepository := mocks.ProductRepository{}
			mockRepository.On(
				"GetAll",
				mock.AnythingOfType("*context.emptyCtx"),
			).Return(tc.mockReturn.products, tc.mockReturn.err)

			service := NewProductService(&mockRepository)
			got, err := service.GetProducts(tc.ctx)

			if tc.errorExpected != nil {
				assert.EqualError(t, err, tc.errorExpected.Error())
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, tc.responseExpected, got)
			}
		})
	}
}
