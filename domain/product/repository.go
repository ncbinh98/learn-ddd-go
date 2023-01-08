package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ncbinh98/learn-ddd-go/aggregate"
)

var (
	ErrProductNotFound = errors.New("no such product")
	ErrProductExists   = errors.New("product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
