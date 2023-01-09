package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound     = errors.New("the customer was not found in repository")
	ErrFailedToAddCustomer  = errors.New("failed to add customer")
	ErrfailedUpdateCustomer = errors.New("failed to update customer")
)

type Repository interface {
	Get(uuid.UUID) (Customer, error)
	Add(Customer) error
	Update(Customer) error
}
