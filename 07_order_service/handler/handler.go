package handler

import (
	"encoding/json"
	"net/http"

	"gochallenges/07_order_service/service"
)

type Handler struct {
	svc service.OrderService
}

func New(svc service.OrderService) *Handler {
	return &Handler{svc: svc}
}

// Register wires routes onto mux.
func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/orders", h.handleOrders)
	mux.HandleFunc("/orders/", h.handleOrderByID)
}

func (h *Handler) handleOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createOrder(w, r)
	case http.MethodGet:
		h.listOrders(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *Handler) handleOrderByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	// TODO: extract id from path, call svc.GetByID, return 404 when not found
}

func (h *Handler) createOrder(w http.ResponseWriter, r *http.Request) {
	// TODO: decode body, call svc.Create, return 400 on validation errors, 201 on success
}

func (h *Handler) listOrders(w http.ResponseWriter, r *http.Request) {
	// TODO: read optional ?status= query param, call svc.List, return 200 with JSON array
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
