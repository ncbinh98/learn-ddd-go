package billing

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("no such Billing")
	ErrProductExists       = errors.New("Billing already exists")
	ErrBillingInvalidTotal = errors.New("Invalid Total")
)

type Repository interface {
	GetAll() ([]Billing, error)
	GetByID(id uuid.UUID) (Billing, error)
	Add(product Billing) error
	Update(product Billing) error
	Delete(id uuid.UUID) error
}
