package service

//holding the business logic here

import (
	"context"

	"github.com/Pavan-pandya1/Microservices/kitchen/services/common/genproto/orders"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	// store
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
