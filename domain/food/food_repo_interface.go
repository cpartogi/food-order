package food

import (
	"context"
	"food-order/domain/food/model"
)

type FoodRepoInterface interface {
	GetMenus(ctx context.Context, menuTypeId string) (res []model.Menus, err error)
}
