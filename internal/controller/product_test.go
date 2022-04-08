package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	pv "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"products-crud/internal/model"
	"products-crud/internal/service/mocks"
	"products-crud/pkg/errors"
	"products-crud/pkg/validator"
)

type request struct {
	body string
}

func TestCreateProductController(t *testing.T) {
	tt := map[string]struct {
		mockReturn         error
		request            request
		bodyExpected       interface{}
		statusCodeExpected int
	}{
		"unmarshal_error": {
			mockReturn: nil,
			request: request{
				body: `{
					"name": "Macbook 2021",
					"supplierId": "1",
					"categoryId": 1,
					"stock" : 200,
					"price" : 3200.49,
					"discontinued" : false
				}`,
			},
			bodyExpected: errors.ApiResponse{
				Message: "unmarshal error data type, got: string, expected: number in supplierId param",
				Code:    "INVALID_REQUEST",
			},
			statusCodeExpected: http.StatusBadRequest,
		},
		"validation_error": {
			mockReturn: nil,
			request: request{
				body: `{
					"name": "Macbook 2021",
					"stock" : 200,
					"price" : 3200.49,
					"discontinued" : false
				}`,
			},
			bodyExpected: errors.ApiResponse{
				Message: "malformed request, please check the following parameters in the request: [supplierId, categoryId]",
				Code:    "INVALID_REQUEST",
			},
			statusCodeExpected: http.StatusBadRequest,
		},
		"could_not_create_product": {
			mockReturn: fmt.Errorf("persistence failed: %w", errors.ErrFailedToSaveProduct),
			request: request{
				body: `{
					"name": "Macbook 2021",
					"supplierId": 1,
					"categoryId": 1,
					"stock" : 200,
					"price" : 3200.49,
					"discontinued" : false
				}`,
			},
			bodyExpected: errors.ApiResponse{
				Message: "persistence failed: product could not be saved",
				Code:    "INTERNAL_SERVER_ERROR",
			},
			statusCodeExpected: http.StatusInternalServerError,
		},
		"create_product_success": {
			mockReturn: nil,
			request: request{
				body: `{
					"name": "Macbook 2021",
					"supplierId": 1,
					"categoryId": 1,
					"stock" : 200,
					"price" : 3200.49,
					"discontinued" : false
				}`,
			},
			bodyExpected:       nil,
			statusCodeExpected: http.StatusNoContent,
		},
	}

	e := echo.New()
	e.Validator = validator.New(pv.New())
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.request.body))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			ctx := e.NewContext(r, w)

			mockService := mocks.ProductService{}
			mockService.On(
				"CreateProduct",
				mock.AnythingOfType("*context.emptyCtx"),
				mock.AnythingOfType("*model.Product"),
			).Return(tc.mockReturn)

			handler := NewProductsHandler(&mockService)
			err := handler.Create(ctx)
			assert.NoError(t, err)

			gotStatusCode := w.Code
			assert.Equal(t, tc.statusCodeExpected, gotStatusCode)

			if name != "create_product_success" {
				var gotBody errors.ApiResponse
				err = json.NewDecoder(w.Body).Decode(&gotBody)
				assert.NoError(t, err)
				assert.Equal(t, tc.bodyExpected, gotBody)
			}
		})
	}
}

func TestGetProductsController(t *testing.T) {
	products := []*model.ProductResponse{
		{
			Id:           1,
			Name:         "Macbook 2021",
			SupplierId:   1,
			CategoryId:   1,
			Stock:        12,
			Price:        2400.32,
			Discontinued: false,
		},
	}
	testCases := map[string]struct {
		mockReturn struct {
			err      error
			products []*model.ProductResponse
		}
		request            request
		bodyExpected       interface{}
		statusCodeExpected int
	}{
		"could_not_get_products": {
			mockReturn: struct {
				err      error
				products []*model.ProductResponse
			}{
				err:      fmt.Errorf("failed getting: %w", errors.ErrFailedToRetrieveProducts),
				products: nil,
			},
			bodyExpected: errors.ApiResponse{
				Message: "failed getting: products could not be retrieved",
				Code:    "INTERNAL_SERVER_ERROR",
			},
			statusCodeExpected: http.StatusInternalServerError,
		},
		"get_products_success": {
			mockReturn: struct {
				err      error
				products []*model.ProductResponse
			}{
				err:      nil,
				products: products,
			},
			bodyExpected: []*model.ProductResponse{
				{
					Id:           1,
					Name:         "Macbook 2021",
					SupplierId:   1,
					CategoryId:   1,
					Stock:        12,
					Price:        2400.32,
					Discontinued: false,
				},
			},
			statusCodeExpected: http.StatusOK,
		},
	}

	e := echo.New()
	e.Validator = validator.New(pv.New())
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(tc.request.body))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			ctx := e.NewContext(r, w)

			mockService := mocks.ProductService{}
			mockService.On(
				"GetProducts",
				mock.AnythingOfType("*context.emptyCtx"),
			).Return(tc.mockReturn.products, tc.mockReturn.err)

			handler := NewProductsHandler(&mockService)
			err := handler.GetAll(ctx)
			assert.NoError(t, err)

			gotStatusCode := w.Code
			assert.EqualValues(t, tc.statusCodeExpected, gotStatusCode)

			if name != "get_products_success" {
				var gotBody errors.ApiResponse
				err = json.NewDecoder(w.Body).Decode(&gotBody)
				assert.NoError(t, err)
				assert.EqualValues(t, tc.bodyExpected, gotBody)
			} else {
				var gotBody []*model.ProductResponse
				err = json.NewDecoder(w.Body).Decode(&gotBody)
				assert.NoError(t, err)
				assert.EqualValues(t, tc.bodyExpected, gotBody)
			}
		})
	}
}
