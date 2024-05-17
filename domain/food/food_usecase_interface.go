package food

import (
	"context"
	"food-order/domain/food/model"
)

type FoodUsecaseInterface interface {
	GetMenus(ctx context.Context) (res []model.Menus, err error)
}
