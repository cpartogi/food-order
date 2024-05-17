package repo

import (
	"context"
	"database/sql"
	"food-order/domain/food"
	"food-order/domain/food/model"
	"food-order/lib/helper"
	"strconv"

	"github.com/google/uuid"
)

type FoodRepo struct {
	db *sql.DB
}

func NewFoodRepo(db *sql.DB) food.FoodRepoInterface {
	return &FoodRepo{
		db: db,
	}
}

func (r *FoodRepo) GetMenus(ctx context.Context) (res []model.Menus, err error) {

	q := `SELECT a.id, a.menu_name, a.menu_description, a.menu_price, b.menu_type_name FROM menus a, menu_types b WHERE a.menu_type_id = b.id ORDER BY b.menu_type_name `

	var rows *sql.Rows
	rows, err = r.db.QueryContext(ctx, q)

	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var menuData model.Menus
		err = rows.Scan(&menuData.Id, &menuData.MenuName, &menuData.MenuDescription, &menuData.MenuPrice, &menuData.MenuTypeName)

		if err != nil {
			return nil, err
		}

		res = append(res, menuData)
	}

	return res, nil
}

func (r *FoodRepo) AddOrders(ctx context.Context, tableNumber int, req []model.AddOrders) (orderId string, err error) {

	orderId = uuid.New().String()
	for _, ord := range req {
		orderItemId := uuid.New().String()
		query := `INSERT into order_items (id, menu_id, order_id, amount, created_by, created_at) values ('` + orderItemId + `' ,'` + ord.MenuId + `', '` + orderId + `',` + strconv.Itoa(ord.Amount) + ` ,'cashier', now())`

		_, err := r.db.ExecContext(ctx, query)

		if err != nil {
			return "", err
		}
	}

	//insert table orders
	orderNumber := helper.GenerateRandomString(10)
	queryInserOrder := `Insert into orders (id, order_number, table_number, created_by, created_at) values ('` + orderId + `', '` + orderNumber + `', ` + strconv.Itoa(tableNumber) + `, 'cashier', now() )`

	_, err = r.db.ExecContext(ctx, queryInserOrder)

	if err != nil {
		return "", err
	}

	return orderId, nil
}
