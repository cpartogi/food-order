package usecase

import (
	"context"
	"food-order/domain/food/model"
)

func (u *FoodUsecase) AddOrders(ctx context.Context, tableNumber int, req []model.AddOrders) (orderId string, err error) {

	orderId, err = u.foodRepo.AddOrders(ctx, tableNumber, req)

	if err != nil {
		return
	}

	return
}
