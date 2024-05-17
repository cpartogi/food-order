package repo

import (
	"context"
	"database/sql"
	"food-order/domain/food"
	"food-order/domain/food/model"
)

type FoodRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) food.FoodRepoInterface {
	return &FoodRepo{
		db: db,
	}
}

func (r *FoodRepo) GetMenus(ctx context.Context, menuTypeId string) (res []model.Menus, err error) {

	q := `SELECT a.id, a.menu_name, a.menu_description, a.menu_price, b.menu_type_name FROM menus a, menu_types b WHERE a.menu_type_id = b.id AND b.id = $1 `

	var rows *sql.Rows

	defer rows.Close()

	rows, err = r.db.QueryContext(ctx, q, menuTypeId)

	if err != nil {
		return
	}

	for rows.Next() {
		var menuData model.Menus
		err = rows.Scan(&menuData.Id, &menuData.MenuName, &menuData.MenuDescription, &menuData.MenuPrice, &menuData.MenuTypeName)

		if err != nil {
			return nil, err
		}

		res = append(res, menuData)
	}

	return
}
