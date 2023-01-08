package services

import (
	"log"

	"github.com/google/uuid"
)

type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	//OrderService to take orders

	OrderService *OrderService

	//Billing Service

	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}

	return t, nil
}

func WithOrderService(os *OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("\nBill the customer: %0.00f\n", price)

	return nil
}
