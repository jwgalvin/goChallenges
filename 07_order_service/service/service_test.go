package service_test

import (
	"errors"
	"testing"

	order_service "gochallenges/07_order_service"
	"gochallenges/07_order_service/repository"
	"gochallenges/07_order_service/service"
)

// fakeStore lets tests control store behaviour without touching the real impl.
type fakeStore struct {
	saved  []*order_service.Order
	findFn func(id string) (*order_service.Order, error)
}

func (f *fakeStore) Save(o *order_service.Order) error {
	f.saved = append(f.saved, o)
	return nil
}
func (f *fakeStore) FindByID(id string) (*order_service.Order, error) {
	if f.findFn != nil {
		return f.findFn(id)
	}
	return nil, repository.ErrNotFound
}
func (f *fakeStore) FindAll(status string) ([]*order_service.Order, error) {
	return f.saved, nil
}

func TestCreate_ValidOrder(t *testing.T) {
	store := &fakeStore{}
	svc := service.New(store)

	req := order_service.CreateOrderRequest{
		Items: []order_service.Item{{ProductID: "p1", Quantity: 2}},
	}
	order, err := svc.Create(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if order.ID == "" {
		t.Error("expected a non-empty ID")
	}
	if order.Status != order_service.StatusPending {
		t.Errorf("status = %q, want %q", order.Status, order_service.StatusPending)
	}
	if len(store.saved) != 1 {
		t.Errorf("expected 1 saved order, got %d", len(store.saved))
	}
}

func TestCreate_EmptyItems(t *testing.T) {
	svc := service.New(&fakeStore{})
	_, err := svc.Create(order_service.CreateOrderRequest{Items: nil})
	if !errors.Is(err, service.ErrEmptyItems) {
		t.Errorf("err = %v, want ErrEmptyItems", err)
	}
}

func TestCreate_ZeroQuantity(t *testing.T) {
	svc := service.New(&fakeStore{})
	req := order_service.CreateOrderRequest{
		Items: []order_service.Item{{ProductID: "p1", Quantity: 0}},
	}
	_, err := svc.Create(req)
	if !errors.Is(err, service.ErrInvalidQuantity) {
		t.Errorf("err = %v, want ErrInvalidQuantity", err)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	svc := service.New(&fakeStore{})
	_, err := svc.GetByID("missing")
	if err == nil {
		t.Fatal("expected an error for missing ID")
	}
}
