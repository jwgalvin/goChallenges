package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	order_service "gochallenges/07_order_service"
	"gochallenges/07_order_service/handler"
	"gochallenges/07_order_service/repository"
)

// stubService satisfies service.OrderService for handler tests.
type stubService struct {
	createFn func(order_service.CreateOrderRequest) (*order_service.Order, error)
	getFn    func(string) (*order_service.Order, error)
	listFn   func(string) ([]*order_service.Order, error)
}

func (s *stubService) Create(req order_service.CreateOrderRequest) (*order_service.Order, error) {
	if s.createFn != nil {
		return s.createFn(req)
	}
	return nil, nil
}
func (s *stubService) GetByID(id string) (*order_service.Order, error) {
	if s.getFn != nil {
		return s.getFn(id)
	}
	return nil, repository.ErrNotFound
}
func (s *stubService) List(status string) ([]*order_service.Order, error) {
	if s.listFn != nil {
		return s.listFn(status)
	}
	return nil, nil
}

func newTestServer(svc *stubService) *httptest.Server {
	mux := http.NewServeMux()
	handler.New(svc).Register(mux)
	return httptest.NewServer(mux)
}

func TestCreateOrder_201(t *testing.T) {
	svc := &stubService{
		createFn: func(req order_service.CreateOrderRequest) (*order_service.Order, error) {
			return &order_service.Order{ID: "abc123", Status: order_service.StatusPending, Items: req.Items}, nil
		},
	}
	srv := newTestServer(svc)
	defer srv.Close()

	body, _ := json.Marshal(order_service.CreateOrderRequest{
		Items: []order_service.Item{{ProductID: "p1", Quantity: 1}},
	})
	resp, err := http.Post(srv.URL+"/orders", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("status = %d, want 201", resp.StatusCode)
	}
}

func TestCreateOrder_400_OnValidationError(t *testing.T) {
	svc := &stubService{
		createFn: func(req order_service.CreateOrderRequest) (*order_service.Order, error) {
			return nil, errors.New("order must contain at least one item")
		},
	}
	srv := newTestServer(svc)
	defer srv.Close()

	body, _ := json.Marshal(order_service.CreateOrderRequest{Items: nil})
	resp, _ := http.Post(srv.URL+"/orders", "application/json", bytes.NewReader(body))
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", resp.StatusCode)
	}
}

func TestGetOrder_404(t *testing.T) {
	srv := newTestServer(&stubService{})
	defer srv.Close()

	resp, _ := http.Get(srv.URL + "/orders/no-such-id")
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("status = %d, want 404", resp.StatusCode)
	}
}

func TestListOrders_200(t *testing.T) {
	svc := &stubService{
		listFn: func(status string) ([]*order_service.Order, error) {
			return []*order_service.Order{{ID: "1", Status: order_service.StatusPending}}, nil
		},
	}
	srv := newTestServer(svc)
	defer srv.Close()

	resp, _ := http.Get(srv.URL + "/orders")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
}
