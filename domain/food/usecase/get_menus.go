package usecase

import (
	"context"
	"food-order/domain/food/model"
)

func (u *FoodUsecase) GetMenus(ctx context.Context, menuTypeId string) (res []model.Menus, err error) {

	res, err = u.foodRepo.GetMenus(ctx, menuTypeId)

	if err != nil {
		return
	}

	return
}
