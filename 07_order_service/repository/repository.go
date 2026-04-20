package repository

import (
	"errors"
	order_service "gochallenges/07_order_service"
)

var ErrNotFound = errors.New("order not found")

// Store is the interface your service depends on.
// The in-memory implementation lives below; a real SQL one would satisfy the same interface.
type Store interface {
	Save(order *order_service.Order) error
	FindByID(id string) (*order_service.Order, error)
	FindAll(statusFilter string) ([]*order_service.Order, error)
}

// InMemoryStore is a thread-unsafe in-memory store (sufficient for tests).
type InMemoryStore struct {
	// TODO: add fields
}

func NewInMemoryStore() *InMemoryStore {
	// TODO: implement
	return &InMemoryStore{}
}

func (s *InMemoryStore) Save(order *order_service.Order) error {
	// TODO: implement
	return nil
}

func (s *InMemoryStore) FindByID(id string) (*order_service.Order, error) {
	// TODO: implement — return ErrNotFound when missing
	return nil, ErrNotFound
}

func (s *InMemoryStore) FindAll(statusFilter string) ([]*order_service.Order, error) {
	// TODO: implement — empty statusFilter means return all
	return nil, nil
}
