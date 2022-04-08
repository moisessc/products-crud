package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"products-crud/internal/model"
	"products-crud/internal/service"
	"products-crud/pkg/errors"
)

// ProductsHandler struct that contains the usecases for the product entity
type ProductsHandler struct {
	service service.ProductService
}

// ProductRequest struct that represents the product request
type ProductRequest struct {
	Name         string  `json:"name" validate:"required"`
	SupplierId   uint    `json:"supplierId" validate:"required"`
	CategoryId   uint    `json:"categoryId" validate:"required"`
	Stock        uint    `json:"stock" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
	Discontinued bool    `json:"discontinued"`
}

// NewProductsHandler creates a new pointer of ProductsHandler struct
func NewProductsHandler(service service.ProductService) *ProductsHandler {
	return &ProductsHandler{
		service: service,
	}
}

// Create invokes the echo handler to create a new product
func (ph *ProductsHandler) Create(c echo.Context) error {
	var req ProductRequest
	if err := c.Bind(&req); err != nil {
		errResponse := errors.MapError(err, errors.UnmarshallErr)
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	if err := c.Validate(req); err != nil {
		errResponse := errors.MapError(err, errors.ValidationErr)
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	product := model.NewProductWithoutId(req.SupplierId, req.CategoryId, req.Stock, req.Name, req.Price, req.Discontinued)
	err := ph.service.CreateProduct(c.Request().Context(), product)
	if err != nil {
		errResponse := errors.MapError(err, errors.DomainErr)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.NoContent(http.StatusNoContent)
}

// GetAll invokes the echo handler to retrieve all the products
func (ph *ProductsHandler) GetAll(c echo.Context) error {
	products, err := ph.service.GetProducts(c.Request().Context())
	if err != nil {
		errResponse := errors.MapError(err, errors.DomainErr)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusOK, products)
}
