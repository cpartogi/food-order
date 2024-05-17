package request

type OrderData struct {
	MenuId string `json:"menuId"`
	Amount int    `json:"amount"`
}

type AddOrder struct {
	OrderDatas []OrderData `json:"orders"`
}
