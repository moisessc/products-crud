package service

import (
	"context"
	"fmt"

	"products-crud/internal/model"
	"products-crud/internal/repository"
)

// ProductService usecases contracts for the product entity
type ProductService interface {
	// CreateProduct usecase to create a new product
	CreateProduct(ctx context.Context, product *model.Product) error
}

// productService struct that implement the ProductService interface
type productService struct {
	repository repository.ProductRepository
}

// NewProductService creates a new pointer of productService struct
func NewProductService(repository repository.ProductRepository) *productService {
	return &productService{
		repository: repository,
	}
}

// CreateProduct implement the interface ProductService.CreateProduct
func (ps *productService) CreateProduct(ctx context.Context, product *model.Product) error {
	pe := product.ToProductEntity()
	if err := ps.repository.Save(ctx, pe); err != nil {
		return fmt.Errorf("persistence failed: %w", err)
	}
	return nil
}
