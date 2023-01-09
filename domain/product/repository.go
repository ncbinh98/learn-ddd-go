package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound = errors.New("no such product")
	ErrProductExists   = errors.New("product already exists")
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(product Product) error
	Update(product Product) error
	Delete(id uuid.UUID) error
}
