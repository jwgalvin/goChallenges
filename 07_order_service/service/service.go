package service

import (
	"errors"
	order_service "gochallenges/07_order_service"
	"gochallenges/07_order_service/repository"
)

var (
	ErrEmptyItems      = errors.New("order must contain at least one item")
	ErrInvalidQuantity = errors.New("item quantity must be greater than zero")
	ErrEmptyProductID  = errors.New("item product_id must not be empty")
)

// OrderService is the interface your handler depends on.
type OrderService interface {
	Create(req order_service.CreateOrderRequest) (*order_service.Order, error)
	GetByID(id string) (*order_service.Order, error)
	List(statusFilter string) ([]*order_service.Order, error)
}

type orderService struct {
	store repository.Store
}

func New(store repository.Store) OrderService {
	return &orderService{store: store}
}

func (s *orderService) Create(req order_service.CreateOrderRequest) (*order_service.Order, error) {
	// TODO: validate req, assign a new UUID, set Status=pending, CreatedAt=now, save and return
	return nil, nil
}

func (s *orderService) GetByID(id string) (*order_service.Order, error) {
	// TODO: delegate to store; translate ErrNotFound appropriately
	return nil, nil
}

func (s *orderService) List(statusFilter string) ([]*order_service.Order, error) {
	// TODO: delegate to store
	return nil, nil
}
