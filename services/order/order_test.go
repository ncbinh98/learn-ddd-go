package order

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ncbinh98/learn-ddd-go/domain/customer"
	"github.com/ncbinh98/learn-ddd-go/domain/product"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "good thing", 0.99)
	if err != nil {
		t.Fatal(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Wine", "nasty thing", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{beer, peenuts, wine}

}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)
	os, err := NewOderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("Binh")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
