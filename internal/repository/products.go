package repository

import (
	"context"
	"database/sql"
	"log"

	"products-crud/internal/model"
	"products-crud/pkg/errors"
)

const (
	// sqlInsertProduct query to insert a new product in the datasource
	sqlInsertProduct = `INSERT INTO products(name, supplier_id, category_id, stock, price, discontinued)
	VALUES ($1, $2, $3, $4, $5, $6);`
	// sqlGetAllProducts query to retrieve the products from the datasource
	sqlGetAllProducts = `SELECT id, name, supplier_id, category_id, stock, price, discontinued FROM products;`
)

// ProductRepository persistence contracts for the product entity
type ProductRepository interface {
	// Save persists a new product in the datasource
	Save(ctx context.Context, pe *model.ProductEntity) error
}

// pqProductRepository struct that implement the ProductRepository interface
type pqProductRepository struct {
	db *sql.DB
}

// NewPqProductRepository creates a new pointer of pqProductRepository struct
func NewPqProductRepository(db *sql.DB) *pqProductRepository {
	return &pqProductRepository{
		db: db,
	}
}

// Save implement the interface ProductRepository.Save
func (pr *pqProductRepository) Save(ctx context.Context, pe *model.ProductEntity) error {
	stmt, err := pr.db.Prepare(sqlInsertProduct)
	if err != nil {
		log.Printf("could not prepare the statement: %v", err)
		return errors.ErrFailedToSaveProduct
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, pe.Name(), pe.SupplierId(), pe.CategoryId(), pe.Stock(), pe.Price(), pe.Discontinued())
	if err != nil {
		log.Printf("could not insert room: %v", err)
		return errors.ErrFailedToSaveProduct
	}

	return nil
}
