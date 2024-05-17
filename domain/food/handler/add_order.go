package http

import (
	"food-order/domain/food/model"
	"food-order/lib/constant"
	"food-order/lib/utils"
	"food-order/schema/request"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *FoodHandler) AddOrder(c echo.Context) error {
	ctx := c.Request().Context()

	tableNo := c.Param("table_no")
	var req request.AddOrder
	c.Bind(&req)

	intTableNo, _ := strconv.Atoi(tableNo)

	var AddOrder []model.AddOrders

	for _, ord := range req.OrderDatas {
		AddOrder = append(AddOrder, model.AddOrders{
			MenuId: ord.MenuId,
			Amount: ord.Amount,
		})
	}

	orderId, err := h.foodUsecase.AddOrders(ctx, intTableNo, AddOrder)

	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessAddData, orderId)
}
