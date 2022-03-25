// Package aggregates holds aggregates that combines many entities into a full object
package aggregate

import (
	"errors"
	"github.com/elwafa/tavern/entity"
	"github.com/elwafa/tavern/valueobject"
	"github.com/google/uuid"
)

var (
	// ErrInvalidPerson is returned when the person is not valid in the NewCustome factory
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

// Customer is a aggregate that combines all entities needed to represent a customer
type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	person *entity.Person
	// a customer can hold many products
	products []*entity.Item
	// a customer can perform many transactions
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new Customer aggregate
// It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	// Validate that the Name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// Create a new person and generate ID
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}
	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

// GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	c.checkPerson()
	c.person.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	c.checkPerson()
	c.person.Name = name
}

// GetName SetName changes the name of the Customer
func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) checkPerson() {
	if c.person == nil {
		c.person = &entity.Person{}
	}
}
