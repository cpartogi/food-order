package http

import (
	"food-order/domain/food"

	"github.com/labstack/echo/v4"
)

type FoodHandler struct {
	foodUsecase food.FoodUsecaseInterface
}

func NewFoodHander(e *echo.Echo, us food.FoodUsecaseInterface) {
	handler := &FoodHandler{
		foodUsecase: us,
	}

	e.GET("/food/menus", handler.GetMenus)
}
