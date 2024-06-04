package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/meteedev/go_choreography/order/internal/ports"
	"github.com/meteedev/go_choreography/pkg/event"
)

type OrderHandler struct {
	OrderService ports.OrderServicePort
}

// func NewHandler(service ports.ApiPort) *OrderHandler {
// 	//return &OrderHandler{Order: service}
// 	return &OrderHandler{}
// }

func NewHandler(s ports.OrderServicePort) *OrderHandler {
	//return &OrderHandler{Order: service}
	return &OrderHandler{
		OrderService: s,
	}
}

func (o *OrderHandler) CreateOrder(c echo.Context) error {

	req := new(event.OrderCreateEvent)

	if err := c.Bind(req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	ctxMQ, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := o.OrderService.CreateOrder(ctxMQ, *req)

	if err != nil {
		return c.String(http.StatusInternalServerError, "create order failed")
	}

	return c.String(http.StatusOK, "create order")
}
