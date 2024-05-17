package food

import (
	"context"
	"food-order/domain/food/model"
)

type FoodRepoInterface interface {
	GetMenus(ctx context.Context) (res []model.Menus, err error)
	AddOrders(ctx context.Context, tableNumber int, req []model.AddOrders) (orderId string, err error)
}
