package billing

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ncbinh98/learn-ddd-go/domain/customer"
	"github.com/ncbinh98/learn-ddd-go/domain/product"
	"github.com/ncbinh98/learn-ddd-go/entity"
)

var (
	ErrMissingValue = errors.New("missing important values")
)

type Billing struct {
	bill     *entity.Bill
	customer *customer.Customer
	products []*product.Product
}

func New(total float64) (Billing, error) {
	if total < 0 {
		return Billing{}, ErrBillingInvalidTotal
	}

	return Billing{
		bill: &entity.Bill{
			ID:    uuid.New(),
			Total: total,
		},
		customer: nil,
		products: make([]*product.Product, 0),
	}, nil
}
