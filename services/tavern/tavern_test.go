package tavern

import (
	"github.com/elwafa/tavern/domain/customer"
	order2 "github.com/elwafa/tavern/services/order"
	"github.com/google/uuid"
	"testing"
)

func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := order2.init_products(t)

	os, err := order2.NewOrderService(
		order2.WithMemoryCustomerRepository(),
		order2.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := customer.NewCustomer("Percy")
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
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}

func Test_MongoTavern(t *testing.T) {
	// Create OrderService
	products := order2.init_products(t)

	os, err := order2.NewOrderService(
		order2.WithMongoCustomerRepository("mongodb://localhost:27017"),
		order2.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := customer.NewCustomer("Percy")
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
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
