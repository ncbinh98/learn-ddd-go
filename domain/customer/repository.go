package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ncbinh98/learn-ddd-go/aggregate"
)

var (
	ErrCustomerNotFound     = errors.New("the customer was not found in repository")
	ErrFailedToAddCustomer  = errors.New("failed to add customer")
	ErrfailedUpdateCustomer = errors.New("failed to update customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
