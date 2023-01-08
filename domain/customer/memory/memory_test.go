package memory

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/ncbinh98/learn-ddd-go/aggregate"
	"github.com/ncbinh98/learn-ddd-go/domain/customer"
)

func TestMemory_GetCustom(t *testing.T) {
	type TestCases struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("binh")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()
	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []TestCases{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("61e0f206-8797-4338-8e88-18b7d4f47a13"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          cust.GetID(),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
