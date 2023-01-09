package tavern

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ncbinh98/learn-ddd-go/domain/product"
	"github.com/ncbinh98/learn-ddd-go/services/order"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)
	os, err := order.NewOderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb:localhost://27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.AddNewCustomer("Binh")

	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)

	if err != nil {
		t.Fatal(err)
	}

}

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
