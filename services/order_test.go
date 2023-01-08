package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ncbinh98/learn-ddd-go/aggregate"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "good thing", 0.99)
	if err != nil {
		t.Fatal(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "nasty thing", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{beer, peenuts, wine}

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

	cust, err := aggregate.NewCustomer("Binh")
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

	err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
