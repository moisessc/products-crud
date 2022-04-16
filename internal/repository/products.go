package repository

import (
	"context"
	"database/sql"
	e "errors"
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
	// sqlGetProductById query to retrieve some product by id from the datasource
	sqlGetProductById = `SELECT id, name, supplier_id, category_id, stock, price, discontinued FROM products WHERE id = $1;`
	// sqlUpdateProduct query to update some product by id in the datasource
	sqlUpdateProduct = `UPDATE products 
	SET name=$1, supplier_id=$2, category_id=$3, stock=$4, price=$5, discontinued=$6 WHERE id = $7;`
)

// ProductRepository persistence contracts for the product entity
type ProductRepository interface {
	// Save persists a new product in the datasource
	Save(ctx context.Context, pe *model.ProductEntity) error
	// GetAll retrieves all the products in the datasource
	GetAll(ctx context.Context) ([]*model.ProductEntity, error)
	// GetById retrieves a product by id in the datasource
	GetById(ctx context.Context, productId uint64) (*model.ProductEntity, error)
	// Update updates a product by id in the datasource
	Update(ctx context.Context, productId uint64, product *model.ProductEntity) (*model.ProductEntity, error)
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
		log.Printf("could not insert the product: %v", err)
		return errors.ErrFailedToSaveProduct
	}

	return nil
}

// GetAll implement the interface ProductRepository.GetAll
func (pr *pqProductRepository) GetAll(ctx context.Context) ([]*model.ProductEntity, error) {
	mp := make([]*model.ProductEntity, 0)
	rows, err := pr.db.QueryContext(ctx, sqlGetAllProducts)
	if err != nil {
		log.Printf("could not retrieve the product: %v", err)
		return nil, errors.ErrFailedToRetrieveProducts
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var id, supplierId, categoryId, stock uint
		var price float64
		var discontinued bool
		err := rows.Scan(&id, &name, &supplierId, &categoryId, &stock, &price, &discontinued)
		if err != nil {
			log.Printf("could not be scan the product: %v", err)
			return nil, errors.ErrFailedToRetrieveProducts
		}
		mp = append(mp, model.NewProductEntity(id, supplierId, categoryId, stock, name, price, discontinued))
	}
	return mp, nil
}

// GetById implement the interface ProductRepository.GetById
func (pr *pqProductRepository) GetById(ctx context.Context, productId uint64) (*model.ProductEntity, error) {
	var name string
	var id, supplierId, categoryId, stock uint
	var price float64
	var discontinued bool

	if err := pr.db.QueryRowContext(
		ctx,
		sqlGetProductById,
		productId,
	).Scan(&id, &name, &supplierId, &categoryId, &stock, &price, &discontinued); err != nil {
		if e.Is(err, sql.ErrNoRows) {
			return nil, errors.ErrProductNotFound
		} else if err != nil {
			log.Printf("could not retrieve the product: %v", err)
			return nil, errors.ErrFailedToRetrieveProduct
		}
	}
	product := model.NewProductEntity(id, supplierId, categoryId, stock, name, price, discontinued)
	return product, nil
}

// Update implement the interface ProductRepository.Update
func (pr *pqProductRepository) Update(ctx context.Context, productId uint64, product *model.ProductEntity) (*model.ProductEntity, error) {
	stmt, err := pr.db.Prepare(sqlUpdateProduct)
	if err != nil {
		log.Printf("could not prepare the statement: %v", err)
		return nil, errors.ErrFailedToUpdateProduct
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		product.Name(),
		product.SupplierId(),
		product.CategoryId(),
		product.Stock(),
		product.Price(),
		product.Discontinued(),
		productId,
	)
	if err != nil {
		log.Printf("could not update the product: %v", err)
		return nil, errors.ErrFailedToUpdateProduct
	}

	return product, nil
}
