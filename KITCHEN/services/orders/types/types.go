package types

import (
	"context"

	"github.com/Pavan-pandya1/Microservices/KITCHEN/KITCHEN/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrer(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
