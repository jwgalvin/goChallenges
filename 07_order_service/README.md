# Task: Order Service

Build a small HTTP order-management service across three layers.
The skeleton compiles; your job is to fill in the TODOs so all tests pass.

## Requirements

1. `POST /orders`       — create an order (validates items are non-empty, quantity > 0)
2. `GET  /orders/{id}`  — fetch order by ID; return 404 if not found
3. `GET  /orders`       — list all orders, optional `?status=` filter

## Rules
- Each layer communicates through the interfaces defined in `service/contracts.go`.
- The handler must NOT know about the repository.
- The repository must NOT import net/http.
- Return structured JSON errors: `{"error": "message"}`.
- All business logic lives in the service layer.

## Signals being tested
- Interface-driven design
- Separation of concerns
- Proper HTTP status codes
- Input validation with meaningful errors
- Testability (handler tests use a fake service, not the real one)
