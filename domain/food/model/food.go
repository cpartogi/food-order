package model

type Menus struct {
	Id              string
	MenuTypeName    string
	MenuName        string
	MenuDescription string
	MenuPrice       int
}

type AddOrders struct {
	MenuId string
	Amount int
}
