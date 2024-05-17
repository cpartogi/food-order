package usecase

import "food-order/domain/food"

type FoodUsecase struct {
	foodRepo food.FoodRepoInterface
}

func NewFoodUsecase(foodRepo food.FoodRepoInterface) food.FoodUsecaseInterface {
	return &FoodUsecase{
		foodRepo: foodRepo,
	}
}
