package service

//holding the business logic here

import (
	"context"

	"github.com/Pavan-pandya1/Microservices/KITCHEN/KITCHEN/services/common/genproto/orders"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	// store
}

// CreateOrer implements types.OrderService.
func (s *OrderService) CreateOrer(context.Context, *orders.Order) error {
	panic("unimplemented")
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDb = append(ordersDb, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersDb
}
