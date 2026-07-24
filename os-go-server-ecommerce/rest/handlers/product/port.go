package product

import "ecommerce/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(id int) error
	Update(domain.Product) (*domain.Product, error)
}
