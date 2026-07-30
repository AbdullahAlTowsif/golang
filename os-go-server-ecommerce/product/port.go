package product

import (
	"ecommerce/domain"
	productHandler "ecommerce/rest/handlers/product"
)

type Service interface {
	productHandler.Service
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productId int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Delete(productId int) error
	Update(p domain.Product) (*domain.Product, error)
}
