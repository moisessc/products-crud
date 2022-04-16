package controller

import (
	"net/http"
	"strconv"

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

// ProductUpdateRequest struct that represents the product update request
type ProductUpdateRequest struct {
	Name         string  `json:"name"`
	SupplierId   uint    `json:"supplierId"`
	CategoryId   uint    `json:"categoryId"`
	Stock        uint    `json:"stock"`
	Price        float64 `json:"price"`
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
		errResponse, code := errors.MapError(err, errors.UnmarshallErr)
		return c.JSON(code, errResponse)
	}

	if err := c.Validate(req); err != nil {
		errResponse, code := errors.MapError(err, errors.ValidationErr)
		return c.JSON(code, errResponse)
	}

	product := model.NewProductWithoutId(req.SupplierId, req.CategoryId, req.Stock, req.Name, req.Price, req.Discontinued)
	err := ph.service.CreateProduct(c.Request().Context(), product)
	if err != nil {
		errResponse, code := errors.MapError(err, errors.DomainErr)
		return c.JSON(code, errResponse)
	}

	return c.NoContent(http.StatusNoContent)
}

// GetAll invokes the echo handler to retrieve all the products
func (ph *ProductsHandler) GetAll(c echo.Context) error {
	products, err := ph.service.GetProducts(c.Request().Context())
	if err != nil {
		errResponse, code := errors.MapError(err, errors.DomainErr)
		return c.JSON(code, errResponse)
	}

	return c.JSON(http.StatusOK, products)
}

// GetById invokes the echo handler to retrieve one product by id
func (ph *ProductsHandler) GetById(c echo.Context) error {
	productId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errResponse, code := errors.MapError(err, errors.InvalidPathParam)
		return c.JSON(code, errResponse)
	}

	product, err := ph.service.GetProductById(c.Request().Context(), productId)
	if err != nil {
		errResponse, code := errors.MapError(err, errors.DomainErr)
		return c.JSON(code, errResponse)
	}

	return c.JSON(http.StatusOK, product)
}

// UpdateProduct invokes the echo handler to update one product by id
func (ph *ProductsHandler) UpdateProduct(c echo.Context) error {
	productId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errResponse, code := errors.MapError(err, errors.InvalidPathParam)
		return c.JSON(code, errResponse)
	}

	var req ProductUpdateRequest
	if err := c.Bind(&req); err != nil {
		errResponse, code := errors.MapError(err, errors.UnmarshallErr)
		return c.JSON(code, errResponse)
	}

	product := model.NewProductWithoutId(
		req.SupplierId,
		req.CategoryId,
		req.Stock,
		req.Name,
		req.Price,
		req.Discontinued,
	)
	productUpdated, err := ph.service.UpdateProduct(c.Request().Context(), productId, product)
	if err != nil {
		errResponse, code := errors.MapError(err, errors.DomainErr)
		return c.JSON(code, errResponse)
	}

	return c.JSON(http.StatusOK, productUpdated)
}
