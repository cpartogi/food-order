package http

import (
	"food-order/lib/constant"
	"food-order/lib/utils"

	"food-order/schema/response"

	"github.com/labstack/echo/v4"
)

func (h *FoodHandler) GetMenus(c echo.Context) error {
	ctx := c.Request().Context()

	menus, err := h.foodUsecase.GetMenus(ctx)

	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	var menuDatas []response.Menus

	for _, mn := range menus {
		menuDatas = append(menuDatas, response.Menus{
			Id:              mn.Id,
			MenuTypeName:    mn.MenuTypeName,
			MenuName:        mn.MenuName,
			MenuDescription: mn.MenuDescription,
			MenuPrice:       mn.MenuPrice,
		})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, menuDatas)
}
