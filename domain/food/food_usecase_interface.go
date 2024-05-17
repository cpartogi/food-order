package food

import (
	"context"
	"food-order/domain/food/model"
)

type FoodUsecaseInterface interface {
	GetMenus(ctx context.Context, menuTypeId string) (res []model.Menus, err error)
}
