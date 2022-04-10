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
	// GetProducts usecase to retrieve all the products
	GetProducts(ctx context.Context) ([]*model.ProductResponse, error)
	// GetProductById usecase to retrieve a product by id
	GetProductById(ctx context.Context, id uint64) (*model.ProductResponse, error)
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

// GetProducts implement the interface ProductService.GetProducts
func (ps *productService) GetProducts(ctx context.Context) ([]*model.ProductResponse, error) {
	products, err := ps.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting: %w", err)
	}

	pr := make([]*model.ProductResponse, 0)
	for _, v := range products {
		pr = append(pr, v.ToProduct().ToProductResponse())
	}
	return pr, nil
}

func (ps *productService) GetProductById(ctx context.Context, id uint64) (*model.ProductResponse, error) {
	product, err := ps.repository.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed getting: %w", err)
	}
	return product.ToProduct().ToProductResponse(), nil
}
