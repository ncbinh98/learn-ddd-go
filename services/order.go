package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/ncbinh98/learn-ddd-go/aggregate"
	"github.com/ncbinh98/learn-ddd-go/domain/customer"
	"github.com/ncbinh98/learn-ddd-go/domain/customer/memory"
	"github.com/ncbinh98/learn-ddd-go/domain/customer/mongo"
	"github.com/ncbinh98/learn-ddd-go/domain/product"
	prodmem "github.com/ncbinh98/learn-ddd-go/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	// Loop through all the cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}

	}
	return os, nil
}

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// Return a function that matches the orderconfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {

	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}

}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()
		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerId uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	//Fetch the customer
	c, err := o.customers.Get(customerId)
	if err != nil {
		return 0, err
	}

	// Get each product, we need product repository
	var products []aggregate.Product
	var total float64
	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return total, nil
}
