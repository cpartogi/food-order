package usecase

import (
	"context"
	"food-order/domain/food/model"
)

func (u *FoodUsecase) GetMenus(ctx context.Context) (res []model.Menus, err error) {

	res, err = u.foodRepo.GetMenus(ctx)

	if err != nil {
		return
	}

	return
}
