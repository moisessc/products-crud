package model

// Product struct for the product entity
type Product struct {
	id           uint
	name         string
	supplierId   uint
	categoryId   uint
	stock        uint
	price        float64
	discontinued bool
}

// ProductResponse struct that represents the product response
type ProductResponse struct {
	Id           uint    `json:"id"`
	Name         string  `json:"name"`
	SupplierId   uint    `json:"supplierId"`
	CategoryId   uint    `json:"categoryId"`
	Stock        uint    `json:"stock"`
	Price        float64 `json:"price"`
	Discontinued bool    `json:"discontinued"`
}

// ProductEntity type for the persistence layer
type ProductEntity Product

// NewProduct creates a new pointer of Product struct
func NewProduct(id, supplierId, categoryId, stock uint, name string, price float64, discontinued bool) *Product {
	return &Product{
		id:           id,
		name:         name,
		supplierId:   supplierId,
		categoryId:   categoryId,
		stock:        stock,
		price:        price,
		discontinued: discontinued,
	}
}

// NewProductWithoutId creates a new pointer of Product struct with the zero value for the id
func NewProductWithoutId(supplierId, categoryId, stock uint, name string, price float64, discontinued bool) *Product {
	return &Product{
		id:           0,
		name:         name,
		supplierId:   supplierId,
		categoryId:   categoryId,
		stock:        stock,
		price:        price,
		discontinued: discontinued,
	}
}

// NewProduct creates a new pointer of ProductEntity type
func NewProductEntity(id, supplierId, categoryId, stock uint, name string, price float64, discontinued bool) *ProductEntity {
	return &ProductEntity{
		id:           id,
		name:         name,
		supplierId:   supplierId,
		categoryId:   categoryId,
		stock:        stock,
		price:        price,
		discontinued: discontinued,
	}
}

// ToProductEntity transform the Product struct in to ProductEntity for the persistence layer
func (p *Product) ToProductEntity() *ProductEntity {
	return &ProductEntity{
		id:           p.id,
		name:         p.name,
		supplierId:   p.supplierId,
		categoryId:   p.categoryId,
		stock:        p.stock,
		price:        p.price,
		discontinued: p.discontinued,
	}
}

// ToProductResponse transform the Product struct in to ProductResponse for the infraestructure layer
func (p *Product) ToProductResponse() *ProductResponse {
	return &ProductResponse{
		Id:           p.id,
		Name:         p.name,
		SupplierId:   p.supplierId,
		CategoryId:   p.categoryId,
		Stock:        p.stock,
		Price:        p.price,
		Discontinued: p.discontinued,
	}
}

// ToProduct transform the ProductEntity struct in to Product for the service layer
func (pe *ProductEntity) ToProduct() *Product {
	return &Product{
		id:           pe.id,
		name:         pe.name,
		supplierId:   pe.supplierId,
		categoryId:   pe.categoryId,
		stock:        pe.stock,
		price:        pe.price,
		discontinued: pe.discontinued,
	}
}

// Id ProductEntity entity id getter
func (p *ProductEntity) Id() uint { return p.id }

// Name ProductEntity entity name getter
func (p *ProductEntity) Name() string { return p.name }

// SupplierId ProductEntity entity supplierId getter
func (p *ProductEntity) SupplierId() uint { return p.supplierId }

// CategoryId ProductEntity entity categoryId getter
func (p *ProductEntity) CategoryId() uint { return p.categoryId }

// Stock ProductEntity entity stock getter
func (p *ProductEntity) Stock() uint { return p.stock }

// Price ProductEntity entity price getter
func (p *ProductEntity) Price() float64 { return p.price }

// Discontinued ProductEntity entity discontinued getter
func (p *ProductEntity) Discontinued() bool { return p.discontinued }
